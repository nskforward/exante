package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Crossrate struct {
	Pair     string `json:"pair"`
	SymbolID string `json:"symbolId"`
	Rate     string `json:"rate"`
}

func (client *Client) GetCrossrate(fromSymbolID, toSymbolID string) (Crossrate, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/crossrates/%s/%s", client.serverAddr, fromSymbolID, toSymbolID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Crossrate{}, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Crossrate{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Crossrate{}, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return Crossrate{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var crossrate Crossrate
	err = json.Unmarshal(data, &crossrate)
	if err != nil {
		return Crossrate{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return crossrate, nil
}
