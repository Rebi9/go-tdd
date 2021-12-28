package money

type AmountWrapper interface {
	getValue() int
}

type Money struct {
	amount int
}

func (this *Money) equals(amount AmountWrapper) bool {
	return this.getValue() == amount.getValue()
}

func (this *Money) getValue() int {
	return this.amount
}