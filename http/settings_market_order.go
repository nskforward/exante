package http

type SettingsMarketOrder struct {
	Filter
}

func (s *SettingsMarketOrder) ClientTag(tag string) *SettingsMarketOrder {
	s.addString("clientTag", tag)
	return s
}

func (s *SettingsMarketOrder) TakeProfit(price float64) *SettingsMarketOrder {
	s.addFloat("takeProfit", price)
	return s
}

func (s *SettingsMarketOrder) StopLoss(price float64) *SettingsMarketOrder {
	s.addFloat("stopLoss", price)
	return s
}
