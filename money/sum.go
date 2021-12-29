package money

type Sum struct {
	Augend *Money
	Added *Money
}

func NewSum(augend *Money, added *Money) *Sum {
	return &Sum{ Augend: augend, Added: added}
}
