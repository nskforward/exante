package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetHistoricalOrders(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	orders, err := client.GetHistoricalOrders(map[string]string{"limit": "10"})
	if err != nil {
		t.Fatal(err)
	}
	if len(orders) == 0 {
		t.Fatal("orders cannot be empty")
	}
	fmt.Println(orders)
}
