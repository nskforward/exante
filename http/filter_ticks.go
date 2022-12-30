package http

import (
	"strconv"
	"time"
)

type FilterTicks struct {
	Filter
}

func (f *FilterTicks) Limit(size int64) *FilterTicks {
	f.addInt("size", size)
	return f
}

func (f *FilterTicks) UseTrades() *FilterTicks {
	f.addString("type", "trades")
	return f
}

func (f *FilterTicks) DateFrom(date time.Time) *FilterTicks {
	f.addString("from", strconv.FormatInt(date.UnixMilli(), 10))
	return f
}

func (f *FilterTicks) DateTo(date time.Time) *FilterTicks {
	f.addString("to", strconv.FormatInt(date.UnixMilli(), 10))
	return f
}
