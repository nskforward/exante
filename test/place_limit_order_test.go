package test

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
	"testing"
	"time"
)

func TestPlaceLimitOrder(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	orders, err := client.PlaceLimitOrder("BTC.USD", exante.BUY, 1.0, 0.001, nil)
	if err != nil {
		t.Fatal(err)
	}
	for i, order := range orders {
		fmt.Println(i+1, order)
	}
}
