package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetOrder(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	order, err := client.GetOrder("0a3b066b-134d-4d61-9bdf-6ea0d0fae281")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(order)
}
