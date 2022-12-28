package test

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
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
	count := 0
	err = client.GetHistoricalOrders(map[string]string{"limit": "10"}, func(order exante.ResponseOrder) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(count)
}
