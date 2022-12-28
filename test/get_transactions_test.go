package test

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
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
	count := 0
	err = client.GetTransactions(map[string]string{"limit": "100", "order": "DESC"}, func(transaction exante.Transaction) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("transactions:", count)
}
