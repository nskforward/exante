package exante

import "fmt"

func (client *Client) PlaceMarketOrder(symbolID string, side OrderSide, quantity float64, settings SettingsMarketOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType": "market",
		"duration":  "immediate_or_cancel",
		"accountId": client.accountID,
		"symbolId":  symbolID,
		"side":      string(side),
		"quantity":  fmt.Sprintf("%.4f", quantity),
	}

	for k, v := range settings.getMap() {
		order[k] = v
	}

	return client.placeOrder(order)
}
