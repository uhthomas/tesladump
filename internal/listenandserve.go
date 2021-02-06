package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
	xcontext "github.com/uhthomas/tesladump/internal/x/context"
)

func ListenAndServe(ctx context.Context, addr string, h http.Handler) error {
	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      h,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(xcontext.Detach(ctx), time.Minute)
		defer cancel()
		return s.Shutdown(ctx)
	})

	g.Go(func() error {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and serve: %w", err)
		}
		return nil
	})
	return g.Wait()
}
