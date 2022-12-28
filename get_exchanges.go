package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var exchanges []Exchange
	err = json.Unmarshal(data, &exchanges)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return exchanges, nil
}
