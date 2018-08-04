package main

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(n int) {
	d.Amount = 10
}

func main() {

}
