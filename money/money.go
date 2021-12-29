package money

type MoneyAccesser interface {
	getAmount() int
	getUnit() string
}

type Money struct {
	amount int
	unit string
}

func (this *Money) equals(accesser MoneyAccesser) bool {
	return this.getAmount() == accesser.getAmount() && this.getUnit() == accesser.getUnit()
}

func (this *Money) getAmount() int {
	return this.amount
}

func (this *Money) getUnit() string {
	return this.unit
}

func NewDollar(amount int) *Money {
	return &Money{ amount: amount, unit: "Dollar" }
}

func (this *Money) times(multiplier int) *Money {
	return NewDollar(this.amount * multiplier)
}
