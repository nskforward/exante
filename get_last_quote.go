package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (client *Client) GetLastQuote(level QuoteLevel, symbolIDs ...string) (Quote, error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s/last?level=%s", client.serverAddr, strings.Join(symbolIDs, ","), level)

	var q Quote

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return q, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return q, err
	}

	defer client.closeResponse(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&q)

	return q, err
}
