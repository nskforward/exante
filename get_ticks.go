package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Tick struct {
	Timestamp  int64  `json:"timestamp"`
	SymbolID   string `json:"symbolId"`
	TradePrice string `json:"price"`
	TradeSize  string `json:"size"`
	Bid        []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"bid"`
	Ask []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"ask"`
}

func (client *Client) GetTicks(symbolId string, filter map[string]string, f func(tick Tick) bool) error {
	var buf bytes.Buffer
	count := 0
	for k, v := range filter {
		if count > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		count++
	}

	url := fmt.Sprintf("%s/md/3.0/ticks/%s?%s", client.serverAddr, symbolId, buf.String())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return err
	}

	defer client.closeResponse(resp.Body)
	d := json.NewDecoder(resp.Body)

	_, err = d.Token()
	if err != nil {
		return err
	}

	for d.More() {
		var tick Tick
		err := d.Decode(&tick)
		if err != nil {
			return err
		}
		if !f(tick) {
			return nil
		}
	}

	_, err = d.Token()
	return err
}
