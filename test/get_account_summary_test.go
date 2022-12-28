package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetAccountSummary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	summary, err := client.GetAccountSummary("USD")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(summary)
}
