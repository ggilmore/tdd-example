package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := Dollar{Amount: 5}
	five.Times(2)
	if five.Amount != 10 {
		t.Errorf("expected $5*2 = $10, got: %d", five.Amount)
	}
}
