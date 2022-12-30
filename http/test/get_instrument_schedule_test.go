package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrumentSchedule(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	intervals, err := client.GetInstrumentSchedule("INTC.NASDAQ")
	if err != nil {
		t.Fatal(err)
	}
	if len(intervals) == 0 {
		t.Fatal("intervals cannot be empty")
	}
	fmt.Println(intervals)
}
