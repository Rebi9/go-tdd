package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{ Money{ amount: amount, unit: "Dollar" }}
}

func (this *Dollar) times(multiplier int) *Dollar {
	return NewDollar(this.amount * multiplier)
}
