package test

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestGetQuoteStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ch, err := client.GetQuoteStream(ctx, "best_price", "BTC.USD", "INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	for q := range ch {
		log.Println("<--", q)
	}
}
