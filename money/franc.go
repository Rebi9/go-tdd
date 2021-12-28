package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{ Money { amount: amount }}
}

func (this *Franc) times(multiplier int) *Franc {
	return NewFranc(this.amount * multiplier)
}
