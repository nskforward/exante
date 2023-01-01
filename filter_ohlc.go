package exante

import (
	"strconv"
	"time"
)

type FilterOHLC struct {
	Filter
}

func (f *FilterOHLC) Limit(size int64) *FilterOHLC {
	f.addInt("size", size)
	return f
}

func (f *FilterOHLC) UseTrades() *FilterOHLC {
	f.addString("type", "trades")
	return f
}

func (f *FilterOHLC) DateFrom(date time.Time) *FilterOHLC {
	f.addString("from", strconv.FormatInt(date.UnixMilli(), 10))
	return f
}

func (f *FilterOHLC) DateTo(date time.Time) *FilterOHLC {
	f.addString("to", strconv.FormatInt(date.UnixMilli(), 10))
	return f
}
