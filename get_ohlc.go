package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (client *Client) GetOHLC(symbolId string, period time.Duration, filter map[string]string) ([]OHLC, error) {
	seconds := int64(period.Seconds())
	switch seconds {
	case 60, 300, 600, 900, 1800, 3600, 14400, 21600, 86400:
	default:
		return nil, fmt.Errorf("period can be one of [1m, 5m, 10m, 15m, 30m, 1h, 4h, 6h, 24h]")
	}

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

	url := fmt.Sprintf("%s/md/3.0/ohlc/%s/%d?%s", client.serverAddr, symbolId, int64(period.Seconds()), buf.String())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	var candles []OHLC
	err = json.Unmarshal(data, &candles)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return candles, nil
}
