package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (client *Client) GetTicks(symbolId string, size int, trades bool) ([]Tick, error) {
	client.refreshAccessToken()

	candleType := "quotes"
	if trades {
		candleType = "trades"
	}

	url := fmt.Sprintf("%s/md/3.0/ticks/%s?size=%d&type=%s", client.serverAddr, symbolId, size, candleType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var ticks []Tick
	err = json.Unmarshal(data, &ticks)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return ticks, nil
}
