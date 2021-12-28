package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Money{ amount: amount }}
}

func (this *Dollar) times(multiplier int) *Dollar {
	return NewDollar(this.amount * multiplier)
}
