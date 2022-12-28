package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetInstrument(symbolId string) (Instrument, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s", client.serverAddr, symbolId)
	var instrument Instrument

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return instrument, err
	}

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return instrument, err
	}

	defer client.closeResponse(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&instrument)

	return instrument, err
}
