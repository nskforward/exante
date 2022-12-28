package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (client *Client) CancelOrder(orderID string) (ResponseOrder, error) {
	body := struct {
		Action string `json:"action"`
	}{
		Action: "cancel",
	}
	dataReq, err := json.Marshal(body)
	if err != nil {
		return ResponseOrder{}, err
	}

	url := fmt.Sprintf("%s/trade/3.0/orders/%s", client.serverAddr, orderID)

	req, err := http.NewRequest("POST", url, bytes.NewReader(dataReq))
	if err != nil {
		return ResponseOrder{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return ResponseOrder{}, err
	}

	defer client.closeResponse(resp.Body)

	var order ResponseOrder
	err = json.NewDecoder(resp.Body).Decode(&order)
	return order, err
}