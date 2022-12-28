package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetTransactions(filter *FilterTransactions, f func(transaction Transaction) bool) error {
	url := fmt.Sprintf("%s/md/3.0/transactions%s", client.serverAddr, filter.string())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return err
	}

	defer client.closeResponse(resp.Body)
	d := json.NewDecoder(resp.Body)

	_, err = d.Token()
	if err != nil {
		return err
	}

	for d.More() {
		var transaction Transaction
		err := d.Decode(&transaction)
		if err != nil {
			return err
		}
		if !f(transaction) {
			return nil
		}
	}

	_, err = d.Token()
	return err
}
