package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (client *Client) GetLastQuote(level string, symbolIDs ...string) ([]Quote, error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s/last?level=%s", client.serverAddr, strings.Join(symbolIDs, ","), level)

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

	var quotes []Quote
	err = json.Unmarshal(data, &quotes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return quotes, nil
}
