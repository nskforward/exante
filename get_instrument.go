package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) GetInstrument(SymbolId string) (Instrument, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s", client.serverAddr, SymbolId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Instrument{}, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return Instrument{}, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Instrument{}, fmt.Errorf("cannot read response body: %w", err)
	}

	var instrument Instrument
	err = json.Unmarshal(data, &instrument)
	if err != nil {
		return Instrument{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return instrument, nil
}
