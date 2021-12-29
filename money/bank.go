package money

type Bank struct {

}

func NewBank() *Bank {
	return &Bank{}
}

func (this *Bank) reduce(source Expression, to string) *Money {
	return source.reduce(to)
}