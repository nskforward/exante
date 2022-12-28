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
