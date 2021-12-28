package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{ Money { amount: amount }}
}

func (this *Franc) times(multiplier int) *Franc {
	return &Franc{ Money { amount: this.amount * multiplier }}
}
