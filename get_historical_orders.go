package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetHistoricalOrders(filter *FilterHistoricalOrders, f func(order ResponseOrder) bool) error {
	url := fmt.Sprintf("%s/trade/3.0/orders%s", client.serverAddr, filter.string())

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
		var order ResponseOrder
		err := d.Decode(&order)
		if err != nil {
			return err
		}
		if !f(order) {
			return nil
		}
	}

	_, err = d.Token()
	return err
}
