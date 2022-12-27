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

func (client *Client) GetQuoteStream(ctx context.Context, level string, symbolIDs ...string) (chan Quote, error) {
	if len(symbolIDs) == 0 {
		return nil, fmt.Errorf("symbolIDs list must be defined")
	}

	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/feed/%s?level=%s", client.serverAddr, strings.Join(symbolIDs, ","), level)

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

	ch := make(chan Quote, 1)

	go func() {
		defer resp.Body.Close()
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

			if q.Event != "" {
				fmt.Println(q.SymbolID, "quotes:", q.Event)
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
