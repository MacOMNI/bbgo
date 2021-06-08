package binance

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mock_key = "CNlnOO6gAIyeJgCIwDUABGVOn5gdIctLuxx1SGRwwjM7u2xvUKHhPzRsvk3UzYu6"
const mock_secret = "ijvxaJdtdtHtO5uxZPu3EFSWbAJo7Y5UmgFn80aUOa5DFxFLpljl8apHjT42kSTI"

func TestAllSymbols(t *testing.T) {
	e := New(mock_key, mock_secret)
	got, err := e.QueryTickers(context.Background())

	balance, err := e.QueryAccountBalances(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("balance = ", balance)
	assert.NoError(t, err)
	if len(got) <= 1 {
		t.Errorf("Binance Exchange: Attempting to get all symbol tickers, but get 1 or less")
	}

}

func TestSomeSymbols(t *testing.T) {
	e := New("mock_key", "mock_secret")
	got, err := e.QueryTickers(context.Background(), "BTCUSDT", "ETHUSDT")

	assert.NoError(t, err)

	if len(got) != 2 {
		t.Errorf("Binance Exchange: Attempting to get two symbols, but number of tickers do not match")

	}
}

func TestSingleSymbol(t *testing.T) {
	e := New("mock_key", "mock_secret")
	got, err := e.QueryTickers(context.Background(), "BTCUSDT")

	assert.NoError(t, err)

	if len(got) != 1 {
		t.Errorf("Binance Exchange: Attempting to get one symbol, but number of tickers do not match")

	}
}
