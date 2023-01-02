package exante_http

type Quote struct {
	Timestamp int64  `json:"timestamp"`
	SymbolID  string `json:"symbolId"`
	Bid       []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"bid"`
	Ask []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"ask"`
}

func (q Quote) Valid() bool {
	return len(q.Bid) > 0 && len(q.Ask) > 0
}
