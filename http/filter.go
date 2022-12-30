package http

import (
	"bytes"
	"fmt"
	"strconv"
)

type Filter struct {
	values map[string]string
}

func (f *Filter) addString(name, value string) {
	if f.values == nil {
		f.values = make(map[string]string)
	}
	f.values[name] = value
}

func (f *Filter) addInt(name string, value int64) {
	if f.values == nil {
		f.values = make(map[string]string)
	}
	f.values[name] = strconv.FormatInt(value, 10)
}

func (f *Filter) addFloat(name string, value float64) {
	if f.values == nil {
		f.values = make(map[string]string)
	}
	f.values[name] = fmt.Sprintf("%.4f", value)
}

func (f *Filter) String() string {
	var buf bytes.Buffer
	count := 0
	for k, v := range f.values {
		if count > 0 {
			buf.WriteString("&")
		} else {
			buf.WriteString("?")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		count++
	}
	return buf.String()
}

func (f *Filter) getMap() map[string]string {
	return f.values
}
