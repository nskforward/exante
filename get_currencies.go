package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseCurrencies struct {
	Currencies []string `json:"currencies"`
}

func (client *Client) GetCurrencies() (ResponseCurrencies, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/crossrates", client.serverAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseCurrencies{}, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseCurrencies{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseCurrencies{}, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return ResponseCurrencies{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var currencies ResponseCurrencies
	err = json.Unmarshal(data, &currencies)
	if err != nil {
		return ResponseCurrencies{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return currencies, nil
}
