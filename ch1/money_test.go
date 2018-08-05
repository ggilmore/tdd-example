package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)

	product := five.Times(2)

	if !product.Equals(*NewDollar(10)) {
		t.Errorf("expected $5*2 = $10, got: %s", product)
	}

	product = five.Times(3)
	if !product.Equals(*NewDollar(15)) {
		t.Errorf("expected $5*3 = $15, got: %s", product)
	}
}

func Test_Equality(t *testing.T) {

	five := NewDollar(5)
	otherFive := NewDollar(5)
	if !five.Equals(*otherFive) {
		t.Error("$5 != %5")
	}

	six := NewDollar(6)
	if five.Equals(*six) {
		t.Error("$5 == $6")
	}
}
