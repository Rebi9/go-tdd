package money

type Sum struct {
	Augend Expression
	Added Expression
}

func NewSum(augend Expression, added Expression) Sum {
	return Sum{ Augend: augend, Added: added}
}

func (this Sum) reduce(bank *Bank, to string) *Money {
	amount := this.Augend.reduce(bank, to).amount + this.Added.reduce(bank, to).amount
	return NewMoney(amount, to)
}

func (this Sum) plus(added Expression) Expression {
	return NewSum(this, added)
}

func (this Sum) times(multiplier int) Expression {
	return NewSum(this.Augend.times(multiplier), this.Added.times(multiplier))
}