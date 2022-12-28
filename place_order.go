package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OrderSide string

const (
	BUY  OrderSide = "buy"
	SELL OrderSide = "sell"
)

func (client *Client) placeOrder(order map[string]string) ([]ResponseOrder, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/trade/3.0/orders", client.serverAddr)

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)
	var orders []ResponseOrder
	err = json.NewDecoder(resp.Body).Decode(&orders)
	return orders, err
}
