package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante/http"
	"testing"
	"time"
)

func TestGetActiveOrders(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var filter exante.FilterActiveOrders

	orders, err := client.GetActiveOrders(filter.SymbolID("BTC.USD").Limit(30))
	if err != nil {
		t.Fatal(err)
	}
	for i, order := range orders {
		fmt.Println(i+1, order)
	}
}
