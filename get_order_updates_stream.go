package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type updatedResponseOrder struct {
	Event string        `json:"event"`
	Order ResponseOrder `json:"order"`
}

func (client *Client) GetOrderUpdatesStream(ctx context.Context) (chan ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/stream/orders", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan ResponseOrder, 1)

	go func() {
		defer client.closeResponse(resp.Body)
		defer close(ch)

		d := json.NewDecoder(resp.Body)

		for {
			select {

			case <-ctx.Done():
				return

			default:
				var updatedResponseOrder updatedResponseOrder

				err := d.Decode(&updatedResponseOrder)
				if err == io.EOF {
					return
				}
				if err != nil {
					fmt.Println("[error] cannot decode order update:", err)
					return
				}

				if updatedResponseOrder.Event != "order" {
					continue
				}

				select {
				case ch <- updatedResponseOrder.Order:
				default: // skip
				}
			}
		}
	}()

	return ch, nil
}
