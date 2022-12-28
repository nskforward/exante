package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OHLC struct {
	Timestamp int64  `json:"timestamp"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
}

func (client *Client) GetOHLC(symbolId string, period time.Duration, filter *FilterOHLC) ([]OHLC, error) {
	seconds := int64(period.Seconds())
	switch seconds {
	case 60, 300, 600, 900, 1800, 3600, 14400, 21600, 86400:
	default:
		return nil, fmt.Errorf("period can be one of [1m, 5m, 10m, 15m, 30m, 1h, 4h, 6h, 24h]")
	}

	url := fmt.Sprintf("%s/md/3.0/ohlc/%s/%d%s", client.serverAddr, symbolId, int64(period.Seconds()), filter.string())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var candles []OHLC
	err = json.NewDecoder(resp.Body).Decode(&candles)

	return candles, err
}
