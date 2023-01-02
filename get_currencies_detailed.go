package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseCurrenciesDetailed struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	ID             string `json:"id"`
	FractionDigits int    `json:"fractionDigits"`
}

func (client *Client) GetCurrenciesDetailed() ([]Currency, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/currencies", client.serverAddr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var currencies responseCurrenciesDetailed
	err = json.NewDecoder(resp.Body).Decode(&currencies)

	return currencies.Currencies, err
}
