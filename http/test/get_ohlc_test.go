package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante/http"
	"testing"
	"time"
)

func TestGetOHLC(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var filter exante.FilterOHLC

	candles, err := client.GetOHLC("BTC.USD", time.Minute, filter.Limit(30))
	if err != nil {
		t.Fatal(err)
	}
	if len(candles) == 0 {
		t.Fatal("candles cannot be empty")
	}
	fmt.Println("candles:", len(candles))
}
