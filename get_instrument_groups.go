package exante

import (
	"encoding/json"
	"fmt"
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

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var groups []InstrumentGroup
	err = json.NewDecoder(resp.Body).Decode(&groups)

	return groups, err
}
