package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetInstrumentsByType(SymbolType string, f func(instrument Instrument) bool) error {
	url := fmt.Sprintf("%s/md/3.0/types/%s", client.serverAddr, SymbolType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return err
	}

	defer client.closeResponse(resp.Body)
	d := json.NewDecoder(resp.Body)

	_, err = d.Token()
	if err != nil {
		return err
	}

	for d.More() {
		var instrument Instrument
		err := d.Decode(&instrument)
		if err != nil {
			return err
		}
		if !f(instrument) {
			return nil
		}
	}

	_, err = d.Token()
	return err
}
