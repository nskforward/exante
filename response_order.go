package exante

type ResponseOrder struct {
	OrderID               string `json:"orderId"`
	PlaceTime             string `json:"placeTime"`
	CurrentModificationID string `json:"currentModificationId"`
	AccountID             string `json:"accountId"`
	Username              string `json:"username"`
	ClientTag             string `json:"clientTag"`
	OrderState            struct {
		Status     string `json:"status"`
		LastUpdate string `json:"lastUpdate"`
		Reason     string `json:"reason"`
		Fills      []struct {
			Timestamp string `json:"timestamp"`
			Quantity  string `json:"quantity"`
			Price     string `json:"price"`
			Position  int64  `json:"position"`
		} `json:"fills"`
	} `json:"orderState"`
	OrderParameters struct {
		Side           string `json:"side"`
		Quantity       string `json:"quantity"`
		OcoGroup       string `json:"ocoGroup"`
		IfDoneParentID string `json:"ifDoneParentId"`
		Duration       string `json:"duration"`
		OrderType      string `json:"orderType"`
		StopPrice      string `json:"stopPrice"`
		LimitPrice     string `json:"limitPrice"`
		PartQuantity   string `json:"partQuantity"`
		PlaceInterval  string `json:"placeInterval"`
		PriceDistance  string `json:"priceDistance"`
		GttExpiration  string `json:"gttExpiration"`
		SymbolID       string `json:"symbolId"`
	} `json:"orderParameters"`
}
