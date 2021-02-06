package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	c  *http.Client
	mc *mongo.Collection
}

func NewService(ctx context.Context, opts ...Option) (*Service, error) {
	var s Service
	for _, opt := range opts {
		if err := opt(ctx, &s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 40*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var out struct {
			Response []struct {
				ID uint64 `json:"id"`
			} `json:"response"`
		}
		if err := s.dump(ctx, "https://owner-api.teslamotors.com/api/1/vehicles", &out); err != nil {
			return err
		}
		for _, v := range out.Response {
			g.Go(func() error {
				return s.dump(ctx,
					"https://owner-api.teslamotors.com/api/1/vehicles/"+strconv.FormatUint(v.ID, 10)+"/vehicle_data",
					nil,
				)
			})
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("dump: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s Service) dump(ctx context.Context, url string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	res, err := s.c.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer res.Body.Close()
	if out == nil {
		out = &json.RawMessage{}
	}
	var buf bytes.Buffer
	if err := json.NewDecoder(io.TeeReader(res.Body, &buf)).Decode(out); err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	if _, err := s.mc.InsertOne(ctx, bson.D{
		{Key: "url", Value: res.Request.URL.String()},
		{Key: "statusCode", Value: res.StatusCode},
		{Key: "response", Value: bson.Raw(buf.Bytes())},
	}); err != nil {
		return fmt.Errorf("insert one: %w", err)
	}
	return nil
}
