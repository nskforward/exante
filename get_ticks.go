package exante

import (
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

func (client *Client) GetTicks(symbolID string, filter *FilterTicks, f func(tick Tick) bool) error {
	filterQuery := ""
	if filter != nil {
		filterQuery = filter.String()
	}

	url := fmt.Sprintf("%s/md/3.0/ticks/%s%s", client.serverAddr, symbolID, filterQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.executeHTTPRequest(req)
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
