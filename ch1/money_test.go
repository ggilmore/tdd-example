package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := Dollar{Amount: 5}
	product := five.Times(2)

	if product != 10 {
		t.Errorf("expected $5*2 = $10, got: %d", product)
	}

	product := five.Times(3)
	if product != 15 {
		t.Errorf("expected $5*3 = $15, got: %d", product)
	}
}
