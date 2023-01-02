package exante_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (client *Client) CancelOrder(orderID string) (ResponseOrder, error) {
	dataReq := []byte(`{"action":"cancel"}`)
	url := fmt.Sprintf("%s/trade/3.0/orders/%s", client.serverAddr, orderID)

	req, err := http.NewRequest("POST", url, bytes.NewReader(dataReq))
	if err != nil {
		return ResponseOrder{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))

	resp, err := client.executeHTTPRequest(req)
	if err != nil {
		return ResponseOrder{}, err
	}

	defer client.closeResponse(resp.Body)

	var order ResponseOrder
	err = json.NewDecoder(resp.Body).Decode(&order)
	return order, err
}
