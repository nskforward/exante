package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseCurrencies struct {
	Currencies []string `json:"currencies"`
}

func (client *Client) GetCurrencies() ([]string, error) {
	url := fmt.Sprintf("%s/md/3.0/crossrates", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var currencies responseCurrencies
	err = json.NewDecoder(resp.Body).Decode(&currencies)

	return currencies.Currencies, err
}
