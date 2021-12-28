package money

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Amount: 10 }
}

func (dollar *Dollar) times(multiplier int) {
	return
}
