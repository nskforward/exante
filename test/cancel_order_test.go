package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelOrder(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	order, err := client.CancelOrder("f96a2282-4035-4ca3-9130-bbcb366b31e8")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(order)
}
