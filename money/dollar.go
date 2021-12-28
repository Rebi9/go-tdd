package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ amount: amount }
}

func (this *Dollar) times(multiplier int) *Dollar {
	return &Dollar{ amount: this.amount * multiplier}
}

func (this *Dollar) equals(dollar *Dollar) bool {
	return this.amount == dollar.amount
}
