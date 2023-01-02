package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Exchange struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (client *Client) GetExchanges() ([]Exchange, error) {
	url := fmt.Sprintf("%s/md/3.0/exchanges", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var exchanges []Exchange
	err = json.NewDecoder(resp.Body).Decode(&exchanges)

	return exchanges, err
}
