package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dataResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(dataResp))
	}
	var orders []ResponseOrder
	err = json.Unmarshal(dataResp, &orders)
	return orders, err
}
