package main

import (
	"currency/currency"
)

func main() {
	nok := currency.NOK{Ammount: 2500}

	currency.SetExchangeRate(11.71)

	currency.PrintCurrency(&nok)
}
