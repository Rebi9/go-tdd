package money

type AmountWrapper interface {
	getValue() int
	getUnit() string
}

type Money struct {
	amount int
	unit string
}

func (this *Money) equals(amount AmountWrapper) bool {
	return this.getValue() == amount.getValue() && this.getUnit() == amount.getUnit()
}

func (this *Money) getValue() int {
	return this.amount
}

func (this *Money) getUnit() string {
	return this.unit
}