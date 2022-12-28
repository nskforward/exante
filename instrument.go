package exante

type Instrument struct {
	SymbolID           string `json:"symbolId"`
	Ticker             string `json:"ticker"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Group              string `json:"group"`
	UnderlyingSymbolId string `json:"underlyingSymbolId"`
	Exchange           string `json:"exchange"`
	Expiration         int64  `json:"expiration"`
	Country            string `json:"country"`
	SymbolType         string `json:"symbolType"`
	OptionData         struct {
		OptionGroupId string `json:"optionGroupId"`
		StrikePrice   string `json:"strikePrice"`
		OptionRight   string `json:"optionRight"`
	} `json:"optionData"`
	MinPriceIncrement string `json:"minPriceIncrement"`
	Currency          string `json:"currency"`
}
