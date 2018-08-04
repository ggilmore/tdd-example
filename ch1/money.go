package main

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{Amount: d.Amount * n}
}

func main() {

}
