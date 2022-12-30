package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrumentSpecification(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	specification, err := client.GetInstrumentSpecification("INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(specification)
}
