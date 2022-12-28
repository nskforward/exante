package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetLastQuote(level QuoteLevel, symbolIDs ...string) ([]Quote, error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s/last?level=%s", client.serverAddr, strings.Join(symbolIDs, ","), level)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var quotes []Quote
	err = json.NewDecoder(resp.Body).Decode(&quotes)

	return quotes, err
}
