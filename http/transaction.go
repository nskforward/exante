package http

type Transaction struct {
	ID            int    `json:"id"`
	UUID          string `json:"uuid"`
	OperationType string `json:"operationType"`
	Timestamp     int    `json:"timestamp"`
	Sum           string `json:"sum"`
	Asset         string `json:"asset"`
	AccountID     string `json:"accountId"`
	SymbolID      string `json:"symbolId"`
	OrderID       string `json:"orderId"`
	OrderPos      int    `json:"orderPos"`
	ValueDate     string `json:"valueDate"`
}
