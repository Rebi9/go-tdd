package money

import (
	"testing";
	"github.com/stretchr/testify/assert"
)

func TestMulitipulication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.times(2))
	assert.Equal(t, NewDollar(15), five.times(3))
}

func TestEquality(t *testing.T) {
	dollar := NewDollar(5)
	assert.Equal(t, true, dollar.equals(NewDollar(5)))
	assert.Equal(t, false, dollar.equals(NewDollar(6)))
	franc := NewFranc(5)
	assert.Equal(t, false, dollar.equals(franc))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).getCurrency())
	assert.Equal(t, "CHF", NewFranc(1).getCurrency())
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.plus(five)
	bank := NewBank()
	reduced := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnSum(t *testing.T) {
	five := NewDollar(5)
	result := five.plus(five)
	sum := result.(Sum)
	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Added)
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferencentCurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestIdentityRate(t *testing.T) {
	bank := NewBank()
	assert.Equal(t, 1, bank.rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)

	result := bank.reduce(fiveBucks.plus(tenFrancs), "USD")

	assert.Equal(t, NewDollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)

	sum := NewSum(fiveBucks, tenFrancs).plus(fiveBucks)
	result := bank.reduce(sum, "USD")

	assert.Equal(t, NewDollar(15), result)
}

func TestSumTimes(t *testing.T) {
		fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)

	sum := NewSum(fiveBucks, tenFrancs).times(2)
	result := bank.reduce(sum, "USD")

	assert.Equal(t, NewDollar(20), result)
}