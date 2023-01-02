package test

import (
	"context"
	exante "github.com/nskforward/exante_http"
	"log"
	"testing"
	"time"
)

const stepSleep = 1 * time.Second

func TestBaseline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()

	step := "build client"
	client, err := Client(ctx)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("+", step)
	time.Sleep(stepSleep)

	{
		step := "get ticks"
		var filter exante.FilterTicks
		err = client.GetTicks("BTC.USD", filter.Limit(5), func(tick exante.Tick) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get account summary"
		_, err = client.GetAccountSummary("USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get account summary by date"
		_, err = client.GetAccountSummaryByDate("USD", time.Now().AddDate(0, 0, -1))
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get account trade stream"
		scoped, cancel2 := context.WithTimeout(ctx, 5*time.Second)
		ch, err := client.GetAccountTradeStream(scoped)
		if err != nil {
			t.Fatal(err)
		}
		for range ch {

		}
		cancel2()
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get active orders"
		_, err = client.GetActiveOrders(nil)
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get crossrate"
		_, err = client.GetCrossrate("EUR", "USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get detailed currencies"
		_, err = client.GetCurrenciesDetailed()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get currencies"
		_, err = client.GetCurrencies()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get daily changes"
		_, err = client.GetDailyChanges("BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get exchanges"
		_, err = client.GetExchanges()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get historical orders"
		var filter exante.FilterHistoricalOrders
		err = client.GetHistoricalOrders(filter.Limit(10), func(order exante.ResponseOrder) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instrument groups"
		_, err = client.GetInstrumentGroups()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instrument schedule"
		_, err = client.GetInstrumentSchedule("BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instrument specification"
		_, err = client.GetInstrumentSpecification("BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instrument"
		_, err = client.GetInstrument("BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instrument types"
		_, err = client.GetInstrumentTypes()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get all instruments"
		err = client.GetInstrumentsAll(func(instrument exante.Instrument) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instruments by exchange"
		err = client.GetInstrumentsByExchange("NASDAQ", func(instrument exante.Instrument) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instruments by group"
		err = client.GetInstrumentsByGroup("INTC", func(instrument exante.Instrument) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get instruments by type"
		err = client.GetInstrumentsByType("FUTURE", func(instrument exante.Instrument) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get last quote"
		_, err = client.GetLastQuote(exante.BestPrice, "BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get ohlc"
		var filter exante.FilterOHLC
		_, err = client.GetOHLC("BTC.USD", time.Minute, filter.Limit(10))
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get order"
		_, err = client.GetOrder("0a3b066b-134d-4d61-9bdf-6ea0d0fae281")
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get order updates stream"
		scoped, cancel2 := context.WithTimeout(ctx, 5*time.Second)
		ch, err := client.GetOrderUpdatesStream(scoped)
		if err != nil {
			t.Fatal(err)
		}
		for range ch {

		}
		cancel2()
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get quote stream"
		scoped, cancel2 := context.WithTimeout(ctx, 5*time.Second)
		ch, err := client.GetQuoteStream(scoped, exante.BestPrice, "BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		for range ch {

		}
		cancel2()
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get symbol trade stream"
		scoped, cancel2 := context.WithTimeout(ctx, 5*time.Second)
		ch, err := client.GetSymbolTradeStream(scoped, "BTC.USD")
		if err != nil {
			t.Fatal(err)
		}
		for range ch {

		}
		cancel2()
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get transactions"
		var filter exante.FilterTransactions
		err = client.GetTransactions(filter.Limit(10), func(transaction exante.Transaction) bool {
			return true
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "get user accounts"
		_, err = client.GetUserAccounts()
		if err != nil {
			t.Fatal(err)
		}
		log.Println("+", step)
		time.Sleep(stepSleep)
	}

	{
		step := "place limit order"
		orders, err := client.PlaceLimitOrder("BTC.USD", exante.BUY, 1.0, 0.01, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(orders) == 0 {
			t.Fatal("orders len must be greater than 0")
		}
		log.Println("+", step)
		time.Sleep(stepSleep)

		step = "cancel order"
		for _, order := range orders {
			_, err = client.CancelOrder(order.OrderID)
			if err != nil {
				t.Fatal(err)
			}
		}
		log.Println("+", step)
	}
}
