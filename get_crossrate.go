package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return Crossrate{}, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Crossrate{}, fmt.Errorf("cannot read response body: %w", err)
	}

	var crossrate Crossrate
	err = json.Unmarshal(data, &crossrate)
	if err != nil {
		return Crossrate{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return crossrate, nil
}
