package exante_http

import (
	"fmt"
)

func (client *Client) PlaceLimitOrder(symbolID string, side OrderSide, price, quantity float64, settings *SettingsLimitOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType":  "limit",
		"duration":   "good_till_cancel",
		"accountId":  client.accountID,
		"symbolId":   symbolID,
		"side":       string(side),
		"quantity":   fmt.Sprintf("%.4f", quantity),
		"limitPrice": fmt.Sprintf("%.4f", price),
	}

	if settings != nil {
		for k, v := range settings.getMap() {
			order[k] = v
			if k == "gttExpiration" {
				order["duration"] = "good_till_time"
			}
		}
	}

	return client.placeOrder(order)
}
