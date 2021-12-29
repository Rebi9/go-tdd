package money

type MoneyAccesser interface {
	getValue() int
	getUnit() string
}

type Money struct {
	amount int
	unit string
}

func (this *Money) equals(accesser MoneyAccesser) bool {
	return this.getValue() == accesser.getValue() && this.getUnit() == accesser.getUnit()
}

func (this *Money) getValue() int {
	return this.amount
}

func (this *Money) getUnit() string {
	return this.unit
}