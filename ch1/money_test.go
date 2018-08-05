package main

import "testing"

func Test_Multiplication(t *testing.T) {
	five := Dollar(5)
	product := five.Times(2)
	if !product.Equals(Dollar(10)) {
		t.Errorf("expected $5*2 = $10, got: %s", product)
	}

	product = five.Times(3)
	if !product.Equals(Dollar(15)) {
		t.Errorf("expected $5*3 = $15, got: %s", product)
	}
}

func Test_Equality(t *testing.T) {

	fiveDollars := Dollar(5)
	otherFiveDollars := Dollar(5)
	if !fiveDollars.Equals(otherFiveDollars) {
		t.Error("$5 != %5")
	}

	sixDollars := Dollar(6)
	if fiveDollars.Equals(sixDollars) {
		t.Error("$5 == $6")
	}

	fiveFrancs := Franc(5)
	if fiveDollars.Equals(fiveFrancs) {
		t.Error("$5 == ₣5")
	}

}

func Test_FrancMultiplication(t *testing.T) {
	five := Franc(5)

	product := five.Times(2)

	if !product.Equals(Franc(10)) {
		t.Errorf("expected ₣:5*2 = ₣:10, got: %s", product)
	}

	product = five.Times(3)
	if !product.Equals(Franc(15)) {
		t.Errorf("expected ₣:5*3 =₣:15, got: %s", product)
	}
}

func Test_Currency(t *testing.T) {

	if "USD" != Dollar(1).Currency {
		t.Errorf("want: USD, got: %s", Dollar(1).Currency)
	}

	if "CHF" != Franc(1).Currency {
		t.Errorf("want: CHF, got: %s", Franc(1).Currency)
	}
}
