package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AccountTrade struct {
	Event     string `json:"event"`
	OrderID   string `json:"orderId"`
	Timestamp string `json:"timestamp"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	Position  string `json:"position"`
}

func (client *Client) GetAccountTradeStream(ctx context.Context) (chan AccountTrade, error) {
	url := fmt.Sprintf("%s/trade/3.0/stream/trades", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/x-json-stream")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan AccountTrade, 1)

	go func() {
		defer client.closeResponse(resp.Body)

		defer close(ch)

		d := json.NewDecoder(resp.Body)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				var t AccountTrade

				err := d.Decode(&t)
				if err == io.EOF {
					return
				}
				if err != nil {
					fmt.Println("[error] cannot decode trade:", err)
					return
				}

				if t.Event != "trade" {
					continue
				}

				select {
				case ch <- t:
				default: // skip
				}
			}
		}
	}()

	return ch, nil
}
