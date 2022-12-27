package test

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestGetOrderUpdatesStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ch, err := client.GetOrderUpdatesStream(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for q := range ch {
		log.Println("<--", q)
	}
}
