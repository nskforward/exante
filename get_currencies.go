package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var currencies responseCurrencies
	err = json.Unmarshal(data, &currencies)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return currencies.Currencies, nil
}
