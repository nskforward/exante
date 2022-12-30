package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetAccountSummary(currency string) (AccountSummary, error) {
	url := fmt.Sprintf("%s/md/3.0/summary/%s/%s", client.serverAddr, client.accountID, currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AccountSummary{}, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return AccountSummary{}, err
	}

	defer client.closeResponse(resp.Body)

	var summary AccountSummary
	err = json.NewDecoder(resp.Body).Decode(&summary)
	return summary, err
}
