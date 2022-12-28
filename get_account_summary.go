package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) GetAccountSummary(accountID, currency string) (AccountSummary, error) {
	url := fmt.Sprintf("%s/md/3.0/summary/%s/%s", client.serverAddr, accountID, currency)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AccountSummary{}, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return AccountSummary{}, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AccountSummary{}, fmt.Errorf("cannot read response body: %w", err)
	}

	var summary AccountSummary
	err = json.Unmarshal(data, &summary)
	if err != nil {
		return AccountSummary{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return summary, nil
}
