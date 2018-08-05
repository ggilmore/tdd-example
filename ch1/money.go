package main

import "fmt"

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: amount}
}

func (d *Dollar) String() string {
	return fmt.Sprintf("$%d", d.Amount)
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{Amount: d.Amount * n}
}

func (d *Dollar) Equals(other Dollar) bool {
	return d.Amount == other.Amount
}

func main() {

}
