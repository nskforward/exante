package exante

func (client *Client) PlaceMarketOrder(symbolID, side, quantity string, settings SettingsMarketOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType": "market",
		"duration":  "immediate_or_cancel",
		"accountId": client.accountID,
		"symbolId":  symbolID,
		"side":      side,
		"quantity":  quantity,
	}

	for k, v := range settings {
		order[k] = v
	}

	return client.placeOrder(order)
}

type SettingsMarketOrder map[string]string

func (s SettingsMarketOrder) ClientTag(tag string) {
	s["clientTag"] = tag
}

func (s SettingsMarketOrder) TakeProfit(price string) {
	s["takeProfit"] = price
}

func (s SettingsMarketOrder) StopLoss(price string) {
	s["stopLoss"] = price
}
