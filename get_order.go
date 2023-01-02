package exante_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetOrder(orderID string) (ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/orders/%s", client.serverAddr, orderID)
	var order ResponseOrder

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return order, err
	}

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return order, err
	}

	defer client.closeResponse(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&order)

	return order, err
}
