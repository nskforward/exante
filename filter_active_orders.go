package exante_http

type FilterActiveOrders struct {
	Filter
}

func (f *FilterActiveOrders) Limit(size int64) *FilterActiveOrders {
	f.addInt("limit", size)
	return f
}

func (f *FilterActiveOrders) AccountID(accountID string) *FilterActiveOrders {
	f.addString("accountId", accountID)
	return f
}

func (f *FilterActiveOrders) SymbolID(symbolID string) *FilterActiveOrders {
	f.addString("symbolId", symbolID)
	return f
}
