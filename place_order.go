package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (client *Client) placeOrder(order map[string]string) ([]ResponseOrder, error) {
	dataReq, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/trade/3.0/orders", client.serverAddr)

	req, err := http.NewRequest("POST", url, bytes.NewReader(dataReq))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return nil, err
	}

	defer client.closeResponse(resp.Body)

	dataResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orders []ResponseOrder
	err = json.Unmarshal(dataResp, &orders)
	return orders, err
}
