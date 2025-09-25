package tests

import (
	"testing"

	"github.com/mainak55512/ccl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Testring for-each function
func TestForEach(t *testing.T) {
	// Should pass
	input_1 := []int{1, 2, 3, 4, 5}
	var result []int

	ccl.ForEach(input_1, func(element int) {
		result = append(result, element)
	})

	require.Len(t, result, len(input_1), "Result should have the same length as input")
	assert.Equal(t, input_1, result, "Result should contain elements in the same order as input")

	// Should pass
	input_2 := []string{"a", "b", "c"}
	var concatenated string

	ccl.ForEach(input_2, func(element string) {
		concatenated += element
	})

	assert.Equal(t, "abc", concatenated, "Concatenated string should match expected value")
}

// Testing find function
func TestFind(t *testing.T) {
	// Should pass
	input_1 := []int{1, 2, 3, 4, 5}

	target := 3
	result := ccl.Find(input_1, target)
	assert.Equal(t, 2, result, "Should be equal")

	// Should fail
	target = 7
	result = ccl.Find(input_1, target)
	assert.Equal(t, -1, result, "Should be equal")

	// Should pass
	input_2 := []string{"a", "b", "c", "d", "e", "f"}

	targetStr := "e"
	result = ccl.Find(input_2, targetStr)
	assert.Equal(t, 4, result, "Should be equal")
}

// Testing map function
func TestMap(t *testing.T) {
	// Should pass
	input := []int{1, 2, 3, 4, 5}

	result := ccl.Map(input, func(element int) int {
		return element * 2
	})

	require.Len(t, result, len(input), "Result should have the same length as input")
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result, "Result should contain elements multiplied by 2")
}

// Testing filter function
func TestFilter(t *testing.T) {
	// Should pass
	input := []int{-1, 2, -3, 4, -5}

	result := ccl.Filter(input, func(element int) bool {
		if element > 0 {
			return true
		}

		return false
	})

	require.Len(t, result, 2, "Result should have the same length as input")
	assert.Equal(t, []int{2, 4}, result, "Result should contain elements which are positive")
}

// Testing resuce function
func TestReduce(t *testing.T) {
	// Should pass
	input_1 := []int{1, 2, 3, 4, 5}
	sum := ccl.Reduce(input_1, func(acc int, elem int) int {
		return acc + elem
	}, 0)

	require.Equal(t, 15, sum, "sum of input should be 15")

	// Should pass
	input_2 := []string{"go", "lang", "is", "cool"}
	concat := ccl.Reduce(input_2, func(acc string, elem string) string {
		return acc + "-" + elem
	}, "hey")

	assert.Equal(t, "hey-go-lang-is-cool", concat)

	// Should pass
	input_3 := []int{1, 2, 3, 4}
	product := ccl.Reduce(input_3, func(acc int, elem int) int {
		return acc * elem
	}, 1)

	assert.Equal(t, 24, product, "product of input should be 24")
}

// Testing reverse function
func TestReverse(t *testing.T) {
	// Should pass
	input_1 := []int{1, 2, 3, 4, 5}

	result := ccl.Reverse(input_1)
	require.Len(t, result, len(input_1), "Result should have the same length as input")
	assert.Equal(t, []int{5, 4, 3, 2, 1}, result, "Result should contain elements in reversed order")

	// Should pass
	input_2 := []string{"a", "b", "c", "d"}

	resultStr := ccl.Reverse(input_2)
	require.Len(t, resultStr, len(input_2), "Result should have the same length as input")
	assert.Equal(t, []string{"d", "c", "b", "a"}, resultStr, "Result should contain elements in reversed order")
}

// Testing unique function
func TestUnique(t *testing.T) {
	// Should pass
	input_1 := []int{1, 2, 2, 2, 3, 4, 4, 4, 4, 5, 5}
	result := ccl.Unique(input_1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result, "Result should contain only unique elements")

	// Should pass
	input_2 := []string{"a", "a", "a", "b", "c", "c"}
	resultStr := ccl.Unique(input_2)
	assert.Equal(t, []string{"a", "b", "c"}, resultStr, "Result should contain only unique elements")
}

// Testing chunk function
func TestChunk(t *testing.T) {
	// Should pass
	arr := []int{1, 2, 3, 4, 5}
	assert.Equal(t, ccl.Chunk(arr, 2), [][]int{{1, 2}, {3, 4}, {5}}, "Result should be array of arrays")

	// Should pass
	arr_str := []string{"a", "a", "a", "b", "c", "c"}
	assert.Equal(t, ccl.Chunk(arr_str, 2), [][]string{{"a", "a"}, {"a", "b"}, {"c", "c"}}, "Result should be array of arrays")
}

// Testing flatten function
func TestFlatten(t *testing.T) {
	// Should pass
	arr := [][]int{{1, 2}, {3, 4}, {5}}
	assert.Equal(t, ccl.Flatten(arr), []int{1, 2, 3, 4, 5}, "Result should be an one-dimentional array")

	// Should Pass
	arr_str := [][]string{{"a", "a"}, {"a", "b"}, {"c", "c"}}
	assert.Equal(t, ccl.Flatten(arr_str), []string{"a", "a", "a", "b", "c", "c"}, "Result should be an one-dimentional array")
}
