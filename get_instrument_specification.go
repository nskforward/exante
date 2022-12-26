package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type InstrumentSpecification struct {
	Leverage           string `json:"leverage"`
	LotSize            string `json:"lotSize"`
	ContractMultiplier string `json:"contractMultiplier"`
	PriceUnit          string `json:"priceUnit"`
	Units              string `json:"units"`
}

func (client *Client) GetInstrumentSpecification(SymbolID string) (InstrumentSpecification, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/symbols/%s/specification", client.serverAddr, SymbolID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return InstrumentSpecification{}, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return InstrumentSpecification{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return InstrumentSpecification{}, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return InstrumentSpecification{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var specification InstrumentSpecification
	err = json.Unmarshal(data, &specification)
	if err != nil {
		return InstrumentSpecification{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return specification, nil
}
