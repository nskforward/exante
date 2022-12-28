package exante

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetHistoricalOrders(filter map[string]string, f func(order ResponseOrder) bool) error {
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

	url := fmt.Sprintf("%s/trade/3.0/orders?%s", client.serverAddr, buf.String())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.executeHttpRequest(req)
	if err != nil {
		return err
	}

	defer client.closeResponse(resp.Body)
	d := json.NewDecoder(resp.Body)

	_, err = d.Token()
	if err != nil {
		return err
	}

	for d.More() {
		var order ResponseOrder
		err := d.Decode(&order)
		if err != nil {
			return err
		}
		if !f(order) {
			return nil
		}
	}

	_, err = d.Token()
	return err
}
