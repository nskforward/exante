package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseInstrumentSchedule struct {
	Intervals []InstrumentSchedule `json:"intervals"`
}

type InstrumentSchedule struct {
	Name   string `json:"name"`
	Period struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	} `json:"period"`
}

func (client *Client) GetInstrumentSchedule(SymbolID string) ([]InstrumentSchedule, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s/schedule", client.serverAddr, SymbolID)
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

	var schedule responseInstrumentSchedule
	err = json.Unmarshal(data, &schedule)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return schedule.Intervals, nil
}
