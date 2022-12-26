package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetCurrenciesDetailed(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	currencies, err := client.GetCurrenciesDetailed()
	if err != nil {
		t.Fatal(err)
	}
	if len(currencies) == 0 {
		t.Fatal("currencies cannot be empty")
	}
	fmt.Println(currencies)
}
