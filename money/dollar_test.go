package money

import (
	"testing";
	"github.com/stretchr/testify/assert"
)

func TestMulitipulication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	assert.Equal(t, NewDollar(10), product)
	product = five.times(3)
	assert.Equal(t, NewDollar(15), product)
}

func TestEquality(t *testing.T) {
	product := NewDollar(5)
	assert.Equal(t, true, product.equals(NewDollar(5)))
	assert.Equal(t, false, product.equals(NewDollar(6)))
}