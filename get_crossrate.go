package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Crossrate struct {
	Pair     string `json:"pair"`
	SymbolID string `json:"symbolId"`
	Rate     string `json:"rate"`
}

func (client *Client) GetCrossrate(fromSymbolID, toSymbolID string) (Crossrate, error) {
	url := fmt.Sprintf("%s/md/3.0/crossrates/%s/%s", client.serverAddr, fromSymbolID, toSymbolID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Crossrate{}, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return Crossrate{}, err
	}

	defer client.closeResponse(resp.Body)
	var crossrate Crossrate
	err = json.NewDecoder(resp.Body).Decode(&crossrate)

	return crossrate, err
}
