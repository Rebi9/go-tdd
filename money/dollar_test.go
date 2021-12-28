package money

import (
	"testing";
	"github.com/stretchr/testify/assert"
)

func TestMulitipulication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	assert.Equal(t, 10, product.Amount)
	product = five.times(3)
	assert.Equal(t, 15, product.Amount)
}