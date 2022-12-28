package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InstrumentGroup struct {
	Group    string   `json:"group"`
	Name     string   `json:"name"`
	Types    []string `json:"types"`
	Exchange string   `json:"exchange"`
}

func (client *Client) GetInstrumentGroups() ([]InstrumentGroup, error) {
	url := fmt.Sprintf("%s/md/3.0/groups", client.serverAddr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	var groups []InstrumentGroup
	err = json.Unmarshal(data, &groups)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return groups, nil
}
