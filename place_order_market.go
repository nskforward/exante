package exante

type SettingsMarketOrder struct {
	Filter
}

func (s *SettingsMarketOrder) ClientTag(tag string) {
	s.addString("clientTag", tag)
}

func (s *SettingsMarketOrder) TakeProfit(price float64) {
	s.addFloat("takeProfit", price)
}

func (s *SettingsMarketOrder) StopLoss(price float64) {
	s.addFloat("stopLoss", price)
}

func (client *Client) PlaceMarketOrder(symbolID, side, quantity string, settings SettingsMarketOrder) ([]ResponseOrder, error) {
	order := map[string]string{
		"orderType": "market",
		"duration":  "immediate_or_cancel",
		"accountId": client.accountID,
		"symbolId":  symbolID,
		"side":      side,
		"quantity":  quantity,
	}

	for k, v := range settings.getMap() {
		order[k] = v
	}

	return client.placeOrder(order)
}
