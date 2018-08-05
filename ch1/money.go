package main

import (
	"fmt"
)

type Money struct {
	amount   int
	Currency string
}

func (m *Money) String() string {
	return fmt.Sprintf("(%s)%d", m.Currency, m.amount)
}

func (m *Money) Times(n int) *Money {
	return &Money{
		amount:   m.amount * n,
		Currency: m.Currency,
	}
}

func (m *Money) Equals(other *Money) bool {
	return m.Currency == other.Currency && m.amount == other.amount
}

func Dollar(amount int) *Money {
	return &Money{
		amount:   amount,
		Currency: "USD",
	}
}

func Franc(amount int) *Money {
	return &Money{
		amount:   amount,
		Currency: "CHF",
	}
}

func main() {

}
