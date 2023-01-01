package test

import (
	"context"
	"github.com/nskforward/exante"
	"log"
	"testing"
	"time"
)

func TestGetQuoteStream(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ch, err := client.GetQuoteStream(ctx, exante.BestPrice, "BTC.USD", "INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	for q := range ch {
		log.Println("<--", q)
	}
}
