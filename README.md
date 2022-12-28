![GoVer](https://img.shields.io/github/go-mod/go-version/nskforward/exante)
[![GoDoc](https://pkg.go.dev/badge/github.com/nskforward/exante?status.svg)](https://pkg.go.dev/github.com/nskforward/exante?tab=doc)
[![Version](https://img.shields.io/github/v/tag/nskforward/exante)](https://github.com/nskforward/exante/tags)
![Issues](https://img.shields.io/github/issues/nskforward/exante)
[![License](https://img.shields.io/github/license/nskforward/exante)](https://github.com/nskforward/exante/blob/main/LICENSE)

# Exante HTTP API v3

A full-featured HTTP API client for Golang

spec: https://api-live.exante.eu/api-docs/

## Installation
```
go get -u github.com/nskforward/exante
```

## Quick start
```
package main

import (
	"context"
	"fmt"
	"github.com/nskforward/exante"
)

func main() {

	client, err := exante.NewClient(
		context.Background(),
		"some account id",              // provided by your broker
		"https://api-demo.exante.eu",   // provided by your broker
		"some client id",               // provided by your broker
		"some app id",                  // provided by your broker
		"some shared key",              // provided by your broker
	)
	if err != nil {
		panic(err)
	}

	quotes, err := client.GetQuoteStream(context.Background(), exante.BestPrice, "BTC.USD")
	if err != nil {
		panic(err)
	}

	for q := range quotes {
		fmt.Println(q.Bid[0].Price, q.Ask[0].Price)
	}
}
```