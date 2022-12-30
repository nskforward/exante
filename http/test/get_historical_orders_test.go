package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante/http"
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

	var filter exante.FilterHistoricalOrders
	count := 0
	err = client.GetHistoricalOrders(filter.Limit(20), func(order exante.ResponseOrder) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("orders:", count)
}
