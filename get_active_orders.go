package exante

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FilterActiveOrders struct {
	Filter
}

func (f *FilterActiveOrders) Limit(size int64) *FilterActiveOrders {
	f.addInt("limit", size)
	return f
}

func (f *FilterActiveOrders) AccountID(accountID string) *FilterActiveOrders {
	f.addString("accountId", accountID)
	return f
}

func (f *FilterActiveOrders) SymbolID(symbolID string) *FilterActiveOrders {
	f.addString("symbolId", symbolID)
	return f
}

func (client *Client) GetActiveOrders(filter *FilterActiveOrders) ([]ResponseOrder, error) {
	url := fmt.Sprintf("%s/trade/3.0/orders/active%s", client.serverAddr, filter.string())

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
