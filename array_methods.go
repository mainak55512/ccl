package ccl

// ForEach function for Array
func ForEach[T comparable](array []T, callback func(element T)) {
	for _, value := range array {
		callback(value)
	}
}

// Map function for Array
func Map[T comparable](array []T, callback func(element T) T) []T {
	var newArr []T
	ForEach(array, func(e T) {
		newArr = append(newArr, callback(e))
	})
	return newArr
}

// Filter function for Array
func Filter[T comparable](array []T, callback func(element T) bool) []T {
	var newArr []T
	ForEach(array, func(e T) {
		if callback(e) {
			newArr = append(newArr, e)
		}
	})
	return newArr
}

// Reduce function for Array
func Reduce[T, K comparable](array []T, callback func(accumulator K, element T) K, initial K) K {
	var acc = initial
	ForEach(array, func(e T) {
		acc = callback(acc, e)
	})
	return acc
}

// Reverse function for Array
func Reverse[T comparable](array []T) []T {
	var newArr []T
	for i := len(array) - 1; i >= 0; i-- {
		newArr = append(newArr, array[i])
	}
	return newArr
}
