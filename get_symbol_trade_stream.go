package exante_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SymbolTrade struct {
	Timestamp int64  `json:"timestamp"`
	SymbolID  string `json:"symbolId"`
	Event     string `json:"event"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

func (client *Client) GetSymbolTradeStream(ctx context.Context, symbolIDs ...string) (chan SymbolTrade, error) {
	if len(symbolIDs) == 0 {
		return nil, fmt.Errorf("symbolIDs list must be defined")
	}

	url := fmt.Sprintf("%s/md/3.0/feed/trades/%s", client.serverAddr, strings.Join(symbolIDs, ","))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/x-json-stream")

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan SymbolTrade, 1)

	go func() {
		defer client.closeResponse(resp.Body)
		defer close(ch)

		fmt.Println("[debug] start trade stream:", strings.Join(symbolIDs, ","))
		defer fmt.Println("[debug] stop trade stream:", strings.Join(symbolIDs, ","))

		d := json.NewDecoder(resp.Body)

		for {
			var t SymbolTrade

			err := d.Decode(&t)
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("[error] cannot decode trade:", err)
				return
			}

			if t.Event != "" {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case ch <- t:
			default: // skip
			}
		}
	}()

	return ch, nil
}
