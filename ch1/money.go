package main

import "fmt"

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) String() string {
	return fmt.Sprintf("$%d", d.amount)
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{amount: d.amount * n}
}

func (d *Dollar) Equals(other Dollar) bool {
	return d.amount == other.amount
}

func main() {

}
