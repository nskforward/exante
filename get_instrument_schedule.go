package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/md/3.0/symbols/%s/schedule", client.serverAddr, SymbolID)
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

	var schedule responseInstrumentSchedule
	err = json.Unmarshal(data, &schedule)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return schedule.Intervals, nil
}
