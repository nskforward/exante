package exante

import (
	"encoding/json"
	"fmt"
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

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var accounts []UserAccount
	err = json.NewDecoder(resp.Body).Decode(&accounts)

	return accounts, err
}
