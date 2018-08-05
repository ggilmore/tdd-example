package main

import (
	"fmt"
	"reflect"
)

type Money interface {
	amount() int
	String() string
}

func Equals(a, b Money) bool {
	return reflect.TypeOf(a).Elem().Name() == reflect.TypeOf(b).Elem().Name() && a.amount() == b.amount()
}

type Dollar struct {
	value int
}

func (d *Dollar) amount() int {
	return d.value
}

func NewDollar(value int) *Dollar {
	return &Dollar{value: value}
}

func (d *Dollar) String() string {
	return fmt.Sprintf("$%d", d.value)
}

func (d *Dollar) Times(n int) Money {
	return &Dollar{value: d.value * n}
}

type Franc struct {
	value int
}

func NewFranc(amount int) *Franc {
	return &Franc{value: amount}
}

func (f *Franc) amount() int {
	return f.value
}

func (f *Franc) String() string {
	return fmt.Sprintf("₣%d", f.value)
}

func (f *Franc) Times(n int) Money {
	return &Franc{value: f.value * n}
}

func main() {

}
