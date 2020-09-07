package in

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckOneWithValidListAndPresentCheck(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	check := 4
	inList, err := CheckOne(input, check)
	assert.Nil(t, err)
	assert.True(t, inList)
}
