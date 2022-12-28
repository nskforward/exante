package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserAccount struct {
	AccountID string `json:"accountId"`
	Status    string `json:"status"`
}

func (client *Client) GetUserAccounts() ([]UserAccount, error) {
	url := fmt.Sprintf("%s/md/3.0/accounts", client.serverAddr)
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

	var accounts []UserAccount
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return accounts, nil
}
