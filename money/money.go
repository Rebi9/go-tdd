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
	return newMoney(amount, "USD")
}

func NewFranc(amount int) *Money {
	return newMoney(amount, "CHF")
}

func newMoney(amount int, currency string) *Money {
	return &Money{ amount: amount, currency: currency }
}

func (this *Money) times(multiplier int) *Money {
	return &Money{ amount: this.amount * multiplier, currency: this.currency }
}

func (this *Money) toString() string {
	return fmt.Sprintf("%v %v", this.amount, this.currency)
}

func (this *Money) plus(added *Money) *Sum {
	return NewSum(this, added)
}
