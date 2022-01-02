package money

type Sum struct {
	Augend *Money
	Added *Money
}

func NewSum(augend *Money, added *Money) Sum {
	return Sum{ Augend: augend, Added: added}
}

func (this Sum) reduce(bank *Bank, to string) *Money {
	amount := this.Augend.amount + this.Added.amount
	return NewMoney(amount, to)
}