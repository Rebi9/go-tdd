package money

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{ amount: amount }
}

func (this *Franc) times(multiplier int) *Franc {
	return &Franc{ amount: this.amount * multiplier}
}

func (this *Franc) equals(franc *Franc) bool {
	return this.amount == franc.amount
}
