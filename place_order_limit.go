package exante

import (
	"fmt"
	"time"
)

type SettingsLimitOrder struct {
	Filter
}

func (s *SettingsLimitOrder) Expiration(date time.Time) {
	s.addString("gttExpiration", date.UTC().Format("2006-01-02 15:04:05"))
}

func (s *SettingsLimitOrder) ClientTag(tag string) {
	s.addString("clientTag", tag)
}

func (s *SettingsLimitOrder) TakeProfit(price float64) {
	s.addString("takeProfit", fmt.Sprintf("%.4f", price))
}

func (s *SettingsLimitOrder) StopLoss(price float64) {
	s.addString("stopLoss", fmt.Sprintf("%.4f", price))
}

func (client *Client) PlaceLimitOrder(symbolID string, side OrderSide, price, quantity float64, settings SettingsLimitOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType":  "limit",
		"duration":   "good_till_cancel",
		"accountId":  client.accountID,
		"symbolId":   symbolID,
		"side":       string(side),
		"quantity":   fmt.Sprintf("%.4f", quantity),
		"limitPrice": fmt.Sprintf("%.4f", price),
	}

	for k, v := range settings.getMap() {
		order[k] = v
		if k == "gttExpiration" {
			order["duration"] = "good_till_time"
		}
	}

	return client.placeOrder(order)
}
