package test

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
	"testing"
	"time"
)

func TestGetInstrumentsByGroup(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	err = client.GetInstrumentsByGroup("INTC", func(instrument exante.Instrument) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("instruments:", count)
}
