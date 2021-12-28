package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Amount: amount }
}

func (dollar *Dollar) times(multiplier int) {
	dollar.Amount *= multiplier
}
