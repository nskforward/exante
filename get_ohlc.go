package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

type FilterOHLC struct {
	Filter
}

func (f *FilterOHLC) Limit(size int64) *FilterOHLC {
	f.addInt("size", size)
	return f
}

func (f *FilterOHLC) UseTrades() *FilterOHLC {
	f.addString("type", "trades")
	return f
}

func (f *FilterOHLC) DateFrom(date time.Time) *FilterOHLC {
	f.addString("from", strconv.FormatInt(date.UnixMilli(), 10))
	return f
}

func (f *FilterOHLC) DateTo(date time.Time) *FilterOHLC {
	f.addString("to", strconv.FormatInt(date.UnixMilli(), 10))
	return f
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
