package exante

import "time"

type FilterHistoricalOrders struct {
	Filter
}

func (f *FilterHistoricalOrders) Limit(size int64) *FilterHistoricalOrders {
	f.addInt("limit", size)
	return f
}

func (f *FilterHistoricalOrders) AccountID(accountID string) *FilterHistoricalOrders {
	f.addString("accountId", accountID)
	return f
}

func (f *FilterHistoricalOrders) DateFrom(date time.Time) *FilterHistoricalOrders {
	f.addString("from", date.UTC().Format("2006-01-02T15:04:05.000Z"))
	return f
}

func (f *FilterHistoricalOrders) DateTo(date time.Time) *FilterHistoricalOrders {
	f.addString("to", date.UTC().Format("2006-01-02T15:04:05.000Z"))
	return f
}
