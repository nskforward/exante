package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrumentTypes(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	types, err := client.GetInstrumentTypes()
	if err != nil {
		t.Fatal(err)
	}
	if len(types) == 0 {
		t.Fatal("types cannot be empty")
	}
	fmt.Println(types)
}
