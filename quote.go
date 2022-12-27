package exante

type Quote struct {
	Timestamp int64  `json:"timestamp"`
	SymbolID  string `json:"symbolId"`
	Event     string `json:"event"`
	Bid       []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"bid"`
	Ask []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"ask"`
}
