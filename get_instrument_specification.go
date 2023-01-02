package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InstrumentSpecification struct {
	Leverage           string `json:"leverage"`
	LotSize            string `json:"lotSize"`
	ContractMultiplier string `json:"contractMultiplier"`
	PriceUnit          string `json:"priceUnit"`
	Units              string `json:"units"`
}

func (client *Client) GetInstrumentSpecification(symbolID string) (InstrumentSpecification, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s/specification", client.serverAddr, symbolID)
	var specification InstrumentSpecification

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return specification, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return specification, err
	}

	defer client.closeResponse(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&specification)

	return specification, err
}
