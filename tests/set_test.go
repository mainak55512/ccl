package tests

import (
	"testing"

	"github.com/mainak55512/ccl"
	"github.com/stretchr/testify/assert"
)

// Testing set to array transformation function
func TestCreateSetFromArray(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1}
	set := ccl.CreateSetFromArray(arr)

	assert.Equal(t, 3, len(set.Array()), "If the length is adequate")
	assert.Contains(t, arr, 1)
	assert.Contains(t, arr, 2)
	assert.Contains(t, arr, 3)
}

// Testing union function
func TestUnion(t *testing.T) {
	arr_1 := []int{1, 2, 3, 4}
	arr_2 := []int{4, 5, 1, 2}

	set_1 := ccl.CreateSetFromArray(arr_1)
	set_2 := ccl.CreateSetFromArray(arr_2)

	result := set_1.Union(set_2)

	assert.Equal(t, 5, len(result.Array()), "If the length is adequate")
	assert.Contains(t, result.Array(), 1)
	assert.Contains(t, result.Array(), 2)
	assert.Contains(t, result.Array(), 3)
	assert.Contains(t, result.Array(), 4)
	assert.Contains(t, result.Array(), 5)
}

// Testing intersection function
func TestIntersection(t *testing.T) {
	arr_1 := []int{1, 2, 3, 4}
	arr_2 := []int{4, 5, 1, 2}

	set_1 := ccl.CreateSetFromArray(arr_1)
	set_2 := ccl.CreateSetFromArray(arr_2)

	result := set_1.Intersection(set_2)

	assert.Equal(t, 3, len(result.Array()), "If the length is adequate")
	assert.Contains(t, result.Array(), 1)
	assert.Contains(t, result.Array(), 2)
	assert.Contains(t, result.Array(), 4)
}

// Testing difference function
func TestDifference(t *testing.T) {
	arr_1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr_2 := []int{4, 5, 1, 2}

	set_1 := ccl.CreateSetFromArray(arr_1)
	set_2 := ccl.CreateSetFromArray(arr_2)

	result := set_1.Difference(set_2)

	assert.Equal(t, 3, len(result.Array()), "If the length is adequate")
	assert.Contains(t, result.Array(), 3)
	assert.Contains(t, result.Array(), 6)
	assert.Contains(t, result.Array(), 7)
}
