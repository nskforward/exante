package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InstrumentType struct {
	ID string `json:"id"`
}

func (client *Client) GetInstrumentTypes() ([]InstrumentType, error) {
	url := fmt.Sprintf("%s/md/3.0/types", client.serverAddr)
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

	var types []InstrumentType
	err = json.Unmarshal(data, &types)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return types, nil
}
