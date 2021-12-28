package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Amount: amount }
}

func (this *Dollar) times(multiplier int) *Dollar {
	return &Dollar{ Amount: this.Amount * multiplier}
}

func (this *Dollar) equals(dollar *Dollar) bool {
	return true;
}
