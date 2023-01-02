package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante_http"
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

	var filter exante.FilterTicks

	count := 0
	err = client.GetTicks("BTC.USD", filter.Limit(100), func(tick exante.Tick) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ticks:", count)
}
