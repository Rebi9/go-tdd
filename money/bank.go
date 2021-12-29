package money

type Bank struct {

}

func NewBank() *Bank {
	return &Bank{}
}

func (this *Bank) reduce(source *Sum, to string) *Money {
	amount := source.Augend.amount + source.Added.amount
	return newMoney(amount, to)
}