package exante_http

type AccountSummary struct {
	AccountID          string `json:"accountId"`
	Currency           string `json:"currency"`
	SessionDate        string `json:"sessionDate"`
	Timestamp          int64  `json:"timestamp"`
	NetAssetValue      string `json:"netAssetValue"`
	FreeMoney          string `json:"freeMoney"`
	MoneyUsedForMargin string `json:"moneyUsedForMargin"`
	MarginUtilization  string `json:"marginUtilization"`
	Currencies         []struct {
		Code           string `json:"code"`
		Value          string `json:"value"`
		ConvertedValue string `json:"convertedValue"`
	} `json:"currencies"`
	Positions []struct {
		SymbolID       string `json:"symbolId"`
		SymbolType     string `json:"symbolType"`
		Quantity       string `json:"quantity"`
		Currency       string `json:"currency"`
		Price          string `json:"price"`
		AveragePrice   string `json:"averagePrice"`
		PnL            string `json:"pnl"`
		ConvertedPnL   string `json:"convertedPnl"`
		Value          string `json:"value"`
		ConvertedValue string `json:"convertedValue"`
	} `json:"positions"`
}
