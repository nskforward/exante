package test

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
	"testing"
	"time"
)

func TestGetTicks(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	err = client.GetTicks("BTC.USD", map[string]string{"size": "1000"}, func(tick exante.Tick) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ticks:", count)
}
