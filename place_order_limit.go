package exante

import (
	"time"
)

func (client *Client) PlaceLimitOrder(symbolID, side, price, quantity string, settings SettingsLimitOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType":  "limit",
		"duration":   "good_till_cancel",
		"accountId":  client.accountID,
		"symbolId":   symbolID,
		"side":       side,
		"quantity":   quantity,
		"limitPrice": price,
	}

	for k, v := range settings {
		order[k] = v
		if k == "gttExpiration" {
			order["duration"] = "good_till_time"
		}
	}

	return client.placeOrder(order)
}

type SettingsLimitOrder map[string]string

func (s SettingsLimitOrder) Expiration(date time.Time) {
	s["gttExpiration"] = date.UTC().Format("2006-01-02 15:04:05")
}

func (s SettingsLimitOrder) ClientTag(tag string) {
	s["clientTag"] = tag
}

func (s SettingsLimitOrder) TakeProfit(price string) {
	s["takeProfit"] = price
}

func (s SettingsLimitOrder) StopLoss(price string) {
	s["stopLoss"] = price
}
