package test

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
