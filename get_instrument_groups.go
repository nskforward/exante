package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type InstrumentGroup struct {
	Group    string   `json:"group"`
	Name     string   `json:"name"`
	Types    []string `json:"types"`
	Exchange string   `json:"exchange"`
}

func (client *Client) GetInstrumentGroups() ([]InstrumentGroup, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/groups", client.serverAddr)
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

	var groups []InstrumentGroup
	err = json.Unmarshal(data, &groups)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return groups, nil
}
