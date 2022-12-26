package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetOHLC(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	candles, err := client.GetOHLC("BTC.USD", time.Minute, 3, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(candles) == 0 {
		t.Fatal("candles cannot be empty")
	}
	fmt.Println(candles)
}
