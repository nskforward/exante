package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (client *Client) GetOrder(orderID string) (ResponseOrder, error) {
	client.refreshAccessToken()

	url := fmt.Sprintf("%s/trade/3.0/orders/%s", client.serverAddr, orderID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseOrder{}, err
	}
	req.WithContext(client.ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", client.accessToken}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResponseOrder{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseOrder{}, fmt.Errorf("cannot read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		return ResponseOrder{}, fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	var order ResponseOrder
	err = json.Unmarshal(data, &order)
	if err != nil {
		return ResponseOrder{}, fmt.Errorf("cannot parse response: %w", err)
	}

	return order, nil
}
