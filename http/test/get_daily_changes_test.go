package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetDailyChanges(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	changes, err := client.GetDailyChanges("BTC.USD")
	if err != nil {
		t.Fatal(err)
	}
	if len(changes) == 0 {
		t.Fatal("changes cannot be empty")
	}
	fmt.Println(changes)
}
