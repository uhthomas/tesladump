package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/uhthomas/tesladump/internal"
)

func Main(ctx context.Context) error {
	addr := flag.String("addr", ":80", "listen address")
	oauth2ConfigPath := flag.String("oauth2-config-path", "oauth2_config.json", "Tesla OAuth2 config file")
	oauth2TokenPath := flag.String("oauth2-token-path", "oauth2_token.json", "Tesla OAuth2 token file")
	muri := flag.String("mongo-uri", "", "Mongo URI")
	flag.Parse()

	if *muri == "" {
		return errors.New("mongo URI must be set")
	}

	s, err := internal.NewService(ctx,
		internal.OAuth2(*oauth2ConfigPath, *oauth2TokenPath),
		internal.Mongo(*muri),
	)
	if err != nil {
		return fmt.Errorf("new service: %w", err)
	}
	return internal.ListenAndServe(ctx, *addr, s)
}

func main() {
	if err := Main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
