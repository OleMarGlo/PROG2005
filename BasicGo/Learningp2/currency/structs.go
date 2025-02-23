package currency

import (
	"fmt"
	"strconv"
)

var exchangeRate float64

/*
*
Sets the exchange rate
*/
func SetExchangeRate(exchange float64) {
	exchangeRate = exchange
}

type Currency interface {
	GenerateString() string
}

type NOK struct {
	Ammount float64
}

func (n *NOK) GenerateString() string {
	return "You hold " + strconv.FormatFloat(n.Ammount, 'f', 2, 64) + " NOK" +
		" This corresponds to " + strconv.Itoa(int(n.Ammount/exchangeRate)) + " Euro"
}

type EUR struct {
	Ammount float64
}

func (e *EUR) GenerateString() string {
	return "You hold " + strconv.FormatFloat(e.Ammount, 'f', 2, 64) + " Euro" +
		" This corresponds to " + strconv.Itoa(int(e.Ammount*exchangeRate)) + " NOK"
}

func PrintCurrency(c Currency) {
	fmt.Println(c.GenerateString())
}
