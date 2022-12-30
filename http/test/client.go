package test

import (
	"context"
	exante "github.com/nskforward/exante/http"
	"os"
)

func Client(ctx context.Context) (*exante.Client, error) {
	return exante.NewClient(
		ctx,
		os.Getenv("EXANTE_ACCOUNT_ID"),
		os.Getenv("EXANTE_SERVER_ADDR"),
		os.Getenv("EXANTE_CLIENT_ID"),
		os.Getenv("EXANTE_APP_ID"),
		os.Getenv("EXANTE_SHARED_KEY"),
	)
}
