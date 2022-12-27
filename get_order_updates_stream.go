package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type updatedResponseOrder struct {
	Event string        `json:"event"`
	Order ResponseOrder `json:"order"`
}

func (client *Client) GetOrderUpdatesStream(ctx context.Context) (chan ResponseOrder, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/trade/3.0/stream/orders", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/x-json-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		data, _ := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return nil, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	ch := make(chan ResponseOrder, 1)

	go func() {
		defer resp.Body.Close()
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
