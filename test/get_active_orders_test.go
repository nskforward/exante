package test

import (
	"context"
	"fmt"
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
	orders, err := client.GetActiveOrders("BTC.USD", 10)
	if err != nil {
		t.Fatal(err)
	}
	for i, order := range orders {
		fmt.Println(i+1, order)
	}
}
