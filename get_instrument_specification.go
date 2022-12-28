package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InstrumentSpecification struct {
	Leverage           string `json:"leverage"`
	LotSize            string `json:"lotSize"`
	ContractMultiplier string `json:"contractMultiplier"`
	PriceUnit          string `json:"priceUnit"`
	Units              string `json:"units"`
}

func (client *Client) GetInstrumentSpecification(SymbolID string) (InstrumentSpecification, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s/specification", client.serverAddr, SymbolID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return InstrumentSpecification{}, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return InstrumentSpecification{}, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return InstrumentSpecification{}, fmt.Errorf("cannot read response body: %w", err)
	}

	var specification InstrumentSpecification
	err = json.Unmarshal(data, &specification)
	if err != nil {
		return InstrumentSpecification{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return specification, nil
}
