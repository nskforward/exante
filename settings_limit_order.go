package exante

import (
	"fmt"
	"time"
)

type SettingsLimitOrder struct {
	Filter
}

func (s *SettingsLimitOrder) Expiration(date time.Time) *SettingsLimitOrder {
	s.addString("gttExpiration", date.UTC().Format("2006-01-02 15:04:05"))
	return s
}

func (s *SettingsLimitOrder) ClientTag(tag string) *SettingsLimitOrder {
	s.addString("clientTag", tag)
	return s
}

func (s *SettingsLimitOrder) TakeProfit(price float64) *SettingsLimitOrder {
	s.addString("takeProfit", fmt.Sprintf("%.4f", price))
	return s
}

func (s *SettingsLimitOrder) StopLoss(price float64) *SettingsLimitOrder {
	s.addString("stopLoss", fmt.Sprintf("%.4f", price))
	return s
}
