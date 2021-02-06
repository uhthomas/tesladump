package internal

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongo(ctx context.Context, uri string) (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	return c.Database(u.Path).Collection("responses"), nil
}
