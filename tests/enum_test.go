package tests

import (
	"testing"

	"github.com/mainak55512/ccl"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Testing insertion
func TestEnum_AddsItemsCorrectly(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.Add("APPLE")
	enum.Add("BANANA")
	enum.Add("CHERRY")

	require.Len(t, enum.Items(), 3)

	assert.NotZero(t, enum.Items()["APPLE"])
	assert.NotZero(t, enum.Items()["BANANA"])
	assert.NotZero(t, enum.Items()["CHERRY"])

	id1 := enum.Items()["APPLE"].(int)
	id2 := enum.Items()["BANANA"].(int)
	id3 := enum.Items()["CHERRY"].(int)

	assert.True(t, id1 < id2 && id2 < id3, "IDs should increment for each unique element")
}

// Testing insertion without duplication
func TestEnum_DoesNotAddDuplicate(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.Add("DOG")
	id1 := enum.Items()["DOG"].(int)
	enum.Add("DOG")
	id2 := enum.Items()["DOG"].(int)

	assert.Equal(t, id1, id2, "ID should not change for duplicate element")
	require.Len(t, enum.Items(), 1, "Duplicate adds should not increase item count")
}

// Testing freezing
func TestEnum_FreezePreventsAdding(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.Add("X")
	enum.Freeze()
	enum.Add("Y")

	require.Len(t, enum.Items(), 1, "No new items should be added after freeze")
	_, exists := enum.Items()["Y"]
	assert.False(t, exists, "'Y' should not be present after freeze")
}

// Testing initial id
func TestEnum_InitialID(t *testing.T) {
	enum := ccl.CreateEnum()

	assert.Equal(t, 1, enum.ID(), "Initial ID should be 1")
	enum.Add("A")
	assert.Equal(t, 2, enum.ID())
	enum.Add("B")
	assert.Equal(t, 3, enum.ID())
}

// Testing AddWithoutValue
func TestEnum_AddWithValue(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.AddWithValue("ONE", 100)
	enum.AddWithValue("TWO", 200)

	assert.Equal(t, 100, enum.Items()["ONE"])
	assert.Equal(t, 200, enum.Items()["TWO"])

}

// Testing Variant reverse lookup
func TestEnum_Variant(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.AddWithValue("ONE", 100)
	enum.AddWithValue("TWO", 200)

	name, err := enum.Variant(100)
	require.NoError(t, err)
	assert.Equal(t, "ONE", name)

	name, err = enum.Variant(300)
	assert.Error(t, err)
	assert.Equal(t, "", name)
}

// Testing match
func TestEnum_Match(t *testing.T) {
	enum := ccl.CreateEnum()

	enum.Add("ONE")
	enum.Add("TWO")

	result := enum.Match("ONE", map[string]func() any{
		"ONE": func() any { return "Matched ONE" },
		"TWO": func() any { return "Matched TWO" },
	})

	assert.Equal(t, "Matched ONE", result)
}
