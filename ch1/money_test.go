package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if !Equals(product, NewDollar(10)) {
		t.Errorf("expected $5*2 = $10, got: %s", product)
	}

	product = five.Times(3)
	if !Equals(product, NewDollar(15)) {
		t.Errorf("expected $5*3 = $15, got: %s", product)
	}
}

func Test_Equality(t *testing.T) {

	fiveDollars := NewDollar(5)
	otherFiveDollars := NewDollar(5)
	if !Equals(fiveDollars, otherFiveDollars) {
		t.Error("$5 != %5")
	}

	sixDollars := NewDollar(6)
	if Equals(fiveDollars, sixDollars) {
		t.Error("$5 == $6")
	}

	fiveFrancs := NewFranc(5)
	otherFiveFrancs := NewFranc(5)
	if !Equals(fiveFrancs, otherFiveFrancs) {
		t.Error("$5 != %5")
	}

	sixFrancs := NewFranc(6)
	if Equals(fiveFrancs, sixFrancs) {
		t.Error("$5 == $6")
	}

}

func Test_FrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	product := five.Times(2)

	if !Equals(product, NewFranc(10)) {
		t.Errorf("expected ₣:5*2 = ₣:10, got: %s", product)
	}

	product = five.Times(3)
	if !Equals(product, NewFranc(15)) {
		t.Errorf("expected ₣:5*3 =₣:15, got: %s", product)
	}
}
