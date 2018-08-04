package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := Dollar{Amount: 5}
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
	if !(&Dollar{Amount: 5}).Equals(Dollar{Amount: 5}) {
		t.Error("$5 != %5")
	}

	if (&Dollar{Amount: 5}).Equals(Dollar{Amount: 6}) {
		t.Error("$5 == $6")
	}
}
