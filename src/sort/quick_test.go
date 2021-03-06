package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test quick sort
func TestQuickSort(t *testing.T) {

	result := QuickSort([]int{1, 34, 6, 7, 10})

	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data quick sort error. ")
}
