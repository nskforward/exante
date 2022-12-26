package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type InstrumentType struct {
	ID string `json:"id"`
}

func (client *Client) GetInstrumentTypes() ([]InstrumentType, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/types", client.serverAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var types []InstrumentType
	err = json.Unmarshal(data, &types)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return types, nil
}
