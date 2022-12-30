package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetUserAccounts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	accounts, err := client.GetUserAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) == 0 {
		t.Fatal("accounts cannot be empty")
	}
	fmt.Println(accounts)
}
