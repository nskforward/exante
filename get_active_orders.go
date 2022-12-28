package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) GetActiveOrders(filter map[string]string) ([]ResponseOrder, error) {
	var buf bytes.Buffer
	count := 0
	for k, v := range filter {
		if count > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		count++
	}

	url := fmt.Sprintf("%s/trade/3.0/orders/active?%s", client.serverAddr, buf.String())

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
