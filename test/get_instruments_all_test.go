package test

import (
	"context"
	"fmt"
	exante "github.com/nskforward/exante_http"
	"runtime"
	"testing"
	"time"
)

// 1755 --> 325 --> 274

func TestGetInstrumentsAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	totalBytes := stat.TotalAlloc

	count := 0
	err = client.GetInstrumentsAll(func(instrument exante.Instrument) bool {
		count++
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	runtime.ReadMemStats(&stat)

	fmt.Println((stat.TotalAlloc-totalBytes)/1024/1024, "MB")

	fmt.Println("instruments:", count)
}
