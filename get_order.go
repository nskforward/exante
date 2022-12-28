package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) GetOrder(orderID string) (ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/orders/%s", client.serverAddr, orderID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseOrder{}, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return ResponseOrder{}, err
	}

	defer client.closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseOrder{}, fmt.Errorf("cannot read response body: %w", err)
	}

	var order ResponseOrder
	err = json.Unmarshal(data, &order)
	if err != nil {
		return ResponseOrder{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return order, nil
}
