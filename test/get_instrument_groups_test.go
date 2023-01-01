package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetInstrumentGroups(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	groups, err := client.GetInstrumentGroups()
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) == 0 {
		t.Fatal("groups cannot be empty")
	}
	fmt.Println(groups)
}
