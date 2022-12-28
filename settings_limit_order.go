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
