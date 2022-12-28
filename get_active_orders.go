package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetActiveOrders(filter *FilterActiveOrders) ([]ResponseOrder, error) {
	filterQuery := ""
	if filter != nil {
		filterQuery = filter.String()
	}
	url := fmt.Sprintf("%s/trade/3.0/orders/active%s", client.serverAddr, filterQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)

	var orders []ResponseOrder
	err = json.NewDecoder(resp.Body).Decode(&orders)

	return orders, nil
}
