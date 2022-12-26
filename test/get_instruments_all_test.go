package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrumentsAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	instruments, err := client.GetInstrumentsAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(instruments) == 0 {
		t.Fatal("instruments cannot be empty")
	}
	fmt.Println("instruments:", len(instruments))
}
