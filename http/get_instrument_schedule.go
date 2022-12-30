package http

import (
	"encoding/json"
	"fmt"
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

func (client *Client) GetInstrumentSchedule(symbolID string) ([]InstrumentSchedule, error) {
	url := fmt.Sprintf("%s/md/3.0/symbols/%s/schedule", client.serverAddr, symbolID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var schedule responseInstrumentSchedule
	err = json.NewDecoder(resp.Body).Decode(&schedule)

	return schedule.Intervals, err
}
