package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetCurrencies(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetCurrencies()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Currencies) == 0 {
		t.Fatal("currencies cannot be empty")
	}
	fmt.Println(resp.Currencies)
}
