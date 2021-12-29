package money

type Sum struct {
	Augend *Money
	Added *Money
}

func NewSum(augend *Money, added *Money) Sum {
	return Sum{ Augend: augend, Added: added}
}

func (this Sum) reduce(to string) *Money {
	amount := this.Augend.amount + this.Added.amount
	return newMoney(amount, to)
}