package ccl

// Map function for Array
func MapArr[T comparable](array []T, callback func(element T) T) []T {
	var newArr []T
	for _, value := range array {
		newArr = append(newArr, callback(value))
	}
	return newArr
}

// Filter function for Array
func FilterArr[T comparable](array []T, callback func(element T) bool) []T {
	var newArr []T
	for _, value := range array {
		if callback(value) {
			newArr = append(newArr, value)
		}
	}
	return newArr
}

// Reduce function for Array
func ReduceArr[T, K comparable](array []T, callback func(accumulator K, element T) K, initial K) K {
	var acc = initial
	for _, value := range array {
		acc = callback(acc, value)
	}
	return acc
}
