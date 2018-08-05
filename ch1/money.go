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

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

func (f *Franc) String() string {
	return fmt.Sprintf("₣:%d", f.amount)
}

func (f *Franc) Times(n int) *Franc {
	return &Franc{amount: f.amount * n}
}

func (f *Franc) Equals(other Franc) bool {
	return f.amount == other.amount
}

func main() {

}
