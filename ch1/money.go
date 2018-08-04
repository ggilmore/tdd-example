package main

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{Amount: d.Amount * n}
}

func (d *Dollar) Equals(other Dollar) bool {
	return d.Amount == other.Amount
}

func main() {

}
