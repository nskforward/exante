package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (client *Client) GetOHLC(symbolId string, period time.Duration, size int, trades bool) ([]OHLC, error) {
	client.refreshAccessToken()

	seconds := int64(period.Seconds())
	switch seconds {
	case 60, 300, 600, 900, 1800, 3600, 14400, 21600, 86400:
	default:
		return nil, fmt.Errorf("period can be one of [1m, 5m, 10m, 15m, 30m, 1h, 4h, 6h, 24h]")
	}

	candleType := "quotes"
	if trades {
		candleType = "trades"
	}

	url := fmt.Sprintf("%s/md/3.0/ohlc/%s/%d?size=%d&type=%s", client.serverAddr, symbolId, int64(period.Seconds()), size, candleType)
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

	var candles []OHLC
	err = json.Unmarshal(data, &candles)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return candles, nil
}
