package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetAccountSummaryByDate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	summary, err := client.GetAccountSummaryByDate("PIX0219.007", "USD", time.Now().AddDate(0, -6, 0))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(summary)
}
