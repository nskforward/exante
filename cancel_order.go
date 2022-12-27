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
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseOrder{}, err
	}
	defer resp.Body.Close()
	dataResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseOrder{}, err
	}
	if resp.StatusCode > 399 {
		return ResponseOrder{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(dataResp))
	}
	var order ResponseOrder
	err = json.Unmarshal(dataResp, &order)
	return order, err
}
