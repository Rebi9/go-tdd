package money

import (
	"testing";
	"github.com/stretchr/testify/assert"
)

func TestMulitipulication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	assert.Equal(t, 10, five.Amount)
}