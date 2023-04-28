package test

import (
	"context"
	"testing"
	"time"
)

func TestGetInstrumentsByExchange(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.GetInstrumentsByExchange("NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
}
