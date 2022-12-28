package exante

import (
	"encoding/json"
	"fmt"
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

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var types []InstrumentType
	err = json.NewDecoder(resp.Body).Decode(&types)

	return types, err
}
