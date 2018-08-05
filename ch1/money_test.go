package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)

	if product.Amount != 10 {
		t.Errorf("expected $5*2 = $10, got: %d", product.Amount)
	}

	product = five.Times(3)
	if product.Amount != 15 {
		t.Errorf("expected $5*3 = $15, got: %d", product.Amount)
	}
}

func Test_Equality(t *testing.T) {

	five := NewDollar(5)
	otherFive := Dollar{Amount: 5}
	if !five.Equals(otherFive) {
		t.Error("$5 != %5")
	}

	six := NewDollar(6)
	if five.Equals(six) {
		t.Error("$5 == $6")
	}
}
