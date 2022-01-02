package money

import (
	"fmt"
)

type MoneyAccesser interface {
	getAmount() int
	getCurrency() string
}

type Money struct {
	amount int
	currency string
}

func (this *Money) equals(accesser MoneyAccesser) bool {
	return this.getAmount() == accesser.getAmount() && this.getCurrency() == accesser.getCurrency()
}

func (this *Money) getAmount() int {
	return this.amount
}

func (this *Money) getCurrency() string {
	return this.currency
}

func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}

func NewMoney(amount int, currency string) *Money {
	return &Money{ amount: amount, currency: currency }
}

func (this *Money) times(multiplier int) Expression {
	return &Money{ amount: this.amount * multiplier, currency: this.currency }
}

func (this *Money) toString() string {
	return fmt.Sprintf("%v %v", this.amount, this.currency)
}

func (this *Money) plus(added Expression) Expression {
	return NewSum(this, added)
}

func (this *Money) reduce(bank *Bank, to string) *Money {
	rate := bank.rate(this.currency, to)
	return NewMoney(this.amount / rate, to)
}