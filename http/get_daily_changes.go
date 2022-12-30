package http

import (
	"encoding/json"
	"fmt"
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

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var changes []DailyChange
	err = json.NewDecoder(resp.Body).Decode(&changes)

	return changes, err
}
