package money

type Bank struct {

}

func NewBank() *Bank {
	return &Bank{}
}

func (this *Bank) reduce(source *Sum, to string) *Money {
	return source.reduce(to)
}