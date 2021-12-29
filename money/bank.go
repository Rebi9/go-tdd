package money

type Bank struct {

}

func NewBank() *Bank {
	return &Bank{}
}

func (this *Bank) reduce(source expression, to string) *Money {
	return NewDollar(10)
}