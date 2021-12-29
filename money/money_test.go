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