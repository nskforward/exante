package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) GetActiveOrders(symbolID string, count int) ([]ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/orders/active?limit=%d&accountId=%s&symbolId=%s", client.serverAddr, count, client.accountID, symbolID)
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

	var orders []ResponseOrder
	err = json.Unmarshal(data, &orders)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return orders, nil
}
