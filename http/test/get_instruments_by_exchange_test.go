package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante/http"
	"testing"
	"time"
)

func TestGetInstrumentsByExchange(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	err = client.GetInstrumentsByExchange("NASDAQ", func(instrument exante.Instrument) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("instruments:", count)
}
