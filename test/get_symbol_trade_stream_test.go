package test

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestGetSymbolTradeStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ch, err := client.GetSymbolTradeStream(ctx, "BTC.USD", "INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	for trade := range ch {
		log.Println("<--", trade)
	}
}
