package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetActiveOrders(filter *FilterActiveOrders) ([]ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/orders/active%s", client.serverAddr, filter.string())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)

	var orders []ResponseOrder
	err = json.NewDecoder(resp.Body).Decode(&orders)

	return orders, nil
}
