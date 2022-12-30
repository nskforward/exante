package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type QuoteLevel string

const (
	BestPrice   QuoteLevel = "best_price"
	MarketDepth QuoteLevel = "market_depth"
)

func (client *Client) GetQuoteStream(ctx context.Context, level QuoteLevel, symbolIDs ...string) (chan Quote, error) {
	if len(symbolIDs) == 0 {
		return nil, fmt.Errorf("symbolIDs list must be defined")
	}

	url := fmt.Sprintf("%s/md/3.0/feed/%s?level=%s", client.serverAddr, strings.Join(symbolIDs, ","), level)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/x-json-stream")

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Quote, 1)

	go func() {
		defer client.closeResponse(resp.Body)
		defer close(ch)

		d := json.NewDecoder(resp.Body)

		for {
			var q Quote

			err := d.Decode(&q)
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("[error] cannot decode quote:", err)
				return
			}

			if !q.Valid() {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case ch <- q:
			default: // skip
			}
		}
	}()

	return ch, nil
}
