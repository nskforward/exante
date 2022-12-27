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

	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/feed/trades/%s", client.serverAddr, strings.Join(symbolIDs, ","))

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

	ch := make(chan SymbolTrade, 1)

	go func() {
		defer resp.Body.Close()
		defer close(ch)

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
				fmt.Println(t.SymbolID, "trades:", t.Event)
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
