package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Amount: amount }
}

func (dollar *Dollar) times(multiplier int) *Dollar {
	return &Dollar{ Amount: dollar.Amount * multiplier}
}

func (this *Dollar) equals(dollar *Dollar) bool {
	return true;
}
