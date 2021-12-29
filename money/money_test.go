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
	assert.Equal(t, true, franc.equals(NewFranc(5)))
	assert.Equal(t, false, franc.equals(NewFranc(6)))
	assert.Equal(t, false, dollar.equals(franc))
}

func TestFrancMulitipulication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.times(2))
	assert.Equal(t, NewFranc(15), five.times(3))
}
