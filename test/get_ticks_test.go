package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetTicks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	ticks, err := client.GetTicks("BTC.USD", 3, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(ticks) == 0 {
		t.Fatal("ticks cannot be empty")
	}
	fmt.Println(ticks)
}
