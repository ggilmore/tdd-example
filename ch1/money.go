package main

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(n int) {
	d.Amount = d.Amount * n
}

func main() {

}
