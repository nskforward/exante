package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetLastQuote(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	quotes, err := client.GetLastQuote("market_depth", "BTC.USD")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(quotes)
}
