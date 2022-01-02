package money

type Expression interface {
	reduce(*Bank, string) *Money
	plus(added Expression) Expression
	times(int) Expression
}