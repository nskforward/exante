package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DailyChange struct {
	SymbolID              string `json:"symbolId"`
	LastSessionClosePrice string `json:"lastSessionClosePrice"`
	DailyChange           string `json:"dailyChange"`
}

func (client *Client) GetDailyChanges(symbolID string) ([]DailyChange, error) {
	url := fmt.Sprintf("%s/md/3.0/change/%s", client.serverAddr, symbolID)
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

	var changes []DailyChange
	err = json.Unmarshal(data, &changes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return changes, nil
}
