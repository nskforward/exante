package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetHistoricalOrders(filter *FilterHistoricalOrders, f func(order ResponseOrder) bool) error {
	filterQuery := ""
	if filter != nil {
		filterQuery = filter.String()
	}
	url := fmt.Sprintf("%s/trade/3.0/orders%s", client.serverAddr, filterQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.executeHTTPRequest(req)
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
