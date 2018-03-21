package basic

func Add(l, r int) int {
	return l + r
}

type Adder struct{}

func (a *Adder) Add(l, r int) int {
	return l + r
}

func (a *Adder) AddMulti(n ...int) int {
	var sum int
	for _, m := range n {
		sum += m
	}
	return sum
}
