package exante

import (
	"bytes"
	"time"
)

type OperationType string

const (
	TypeTrade             OperationType = "TRADE"
	TypeCommission        OperationType = "COMMISSION"
	TypeFee               OperationType = "FEE"
	TypeDividend          OperationType = "DIVIDEND"
	TypeCorporateAction   OperationType = "CORPORATE ACTION"
	TypeTax               OperationType = "TAX"
	TypeRollover          OperationType = "ROLLOVER"
	TypeSpecialFee        OperationType = "SPECIAL FEE"
	TypeStockSplit        OperationType = "STOCK SPLIT"
	TypeAutoconversion    OperationType = "AUTOCONVERSION"
	TypeVariationMargin   OperationType = "VARIATION MARGIN"
	TypeExcessMarginFee   OperationType = "EXCESS MARGIN FEE"
	TypeManualCloseOutFee OperationType = "MANUAL CLOSE-OUT FEE"
	TypeFundingWithdrawal OperationType = "FUNDING/WITHDRAWAL"
	TypeExercise          OperationType = "EXERCISE"
	TypeInterest          OperationType = "INTEREST"
	TypeTradeCorrection   OperationType = "TRADE CORRECTION"
)

type FilterTransactions struct {
	Filter
}

func (f *FilterTransactions) UUID(value string) *FilterTransactions {
	f.addString("uuid", value)
	return f
}

func (f *FilterTransactions) AccountID(value string) *FilterTransactions {
	f.addString("accountId", value)
	return f
}

func (f *FilterTransactions) SymbolID(value string) *FilterTransactions {
	f.addString("symbolId", value)
	return f
}

func (f *FilterTransactions) Currency(value string) *FilterTransactions {
	f.addString("asset", value)
	return f
}

func (f *FilterTransactions) OperationTypes(values ...OperationType) *FilterTransactions {
	var buf bytes.Buffer
	for i, item := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(item))
	}

	f.addString("operationType", buf.String())
	return f
}

func (f *FilterTransactions) Offset(value int64) *FilterTransactions {
	f.addInt("offset", value)
	return f
}

func (f *FilterTransactions) Limit(value int64) *FilterTransactions {
	f.addInt("limit", value)
	return f
}

func (f *FilterTransactions) Desc() *FilterTransactions {
	f.addString("order", "DESC")
	return f
}

func (f *FilterTransactions) DateFrom(value time.Time) *FilterTransactions {
	f.addString("fromDate", value.UTC().Format("2006-01-02"))
	return f
}

func (f *FilterTransactions) DateTo(value time.Time) *FilterTransactions {
	f.addString("toDate", value.UTC().Format("2006-01-02"))
	return f
}

func (f *FilterTransactions) OrderID(value string) *FilterTransactions {
	f.addString("orderId", value)
	return f
}

func (f *FilterTransactions) Position(value int64) *FilterTransactions {
	f.addInt("orderPos", value)
	return f
}
