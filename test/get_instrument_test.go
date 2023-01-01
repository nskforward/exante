package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrument(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	instrument, err := client.GetInstrument("INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(instrument)
}
