package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetTransactions(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	transactions, err := client.GetTransactions(map[string]string{"limit": "2", "order": "DESC", "symbolId": "BTC.USD", "operationType": "ROLLOVER"})
	if err != nil {
		t.Fatal(err)
	}
	if len(transactions) == 0 {
		t.Fatal("transactions cannot be empty")
	}
	fmt.Println(transactions)
}
