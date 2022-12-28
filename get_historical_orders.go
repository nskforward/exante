package exante

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FilterHistoricalOrders struct {
	Filter
}

func (f *FilterHistoricalOrders) Limit(size int64) *FilterHistoricalOrders {
	f.addInt("limit", size)
	return f
}

func (f *FilterHistoricalOrders) AccountID(accountID string) *FilterHistoricalOrders {
	f.addString("accountId", accountID)
	return f
}

func (f *FilterHistoricalOrders) DateFrom(date time.Time) *FilterHistoricalOrders {
	f.addString("from", date.UTC().Format("2006-01-02T15:04:05.000Z"))
	return f
}

func (f *FilterHistoricalOrders) DateTo(date time.Time) *FilterHistoricalOrders {
	f.addString("to", date.UTC().Format("2006-01-02T15:04:05.000Z"))
	return f
}

func (client *Client) GetHistoricalOrders(filter *FilterHistoricalOrders, f func(order ResponseOrder) bool) error {
	url := fmt.Sprintf("%s/trade/3.0/orders%s", client.serverAddr, filter.string())
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
