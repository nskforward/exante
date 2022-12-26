package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetCrossrate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	crossrate, err := client.GetCrossrate("EUR", "USD")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(crossrate)
}
