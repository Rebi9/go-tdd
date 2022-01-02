package money

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{ rates: map[Pair]int{} }
}

func (this *Bank) reduce(source Expression, to string) *Money {
	return source.reduce(this, to)
}

func (this *Bank) addRate(from string, to string, rate int) {
	this.rates[NewPair(from, to)] = rate
}

func (this *Bank) rate(from string, to string) int {
	if from == to {
		return 1
	}
	return this.rates[NewPair(from, to)]
}