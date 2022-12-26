package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetExchanges(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	exchanges, err := client.GetExchanges()
	if err != nil {
		t.Fatal(err)
	}
	if len(exchanges) == 0 {
		t.Fatal("exchanges cannot be empty")
	}
	fmt.Println(exchanges)
}
