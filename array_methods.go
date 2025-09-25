package ccl

// ForEach function for Array
/*
Usage:

arr := []int{1,2,3,4,5}
ForEach(arr, func(e int) {
	fmt.Println(e)
})
*/
func ForEach[T any](array []T, callback func(element T)) {
	for _, value := range array {
		callback(value)
	}
}

// Find function for Array
/*
Usage:

arr := []int{1,2,3,4,5}
Find(arr, 4) // => 1 or, -1
*/
func Find[T comparable](array []T, element T) int {
	for i, val := range array {
		if val == element {
			return i
		}
	}
	return -1
}

// Map function for Array
/*
Usage:

arr := []int{1,2,3,4,5}
newArr := Map(arr, func(e int) int {
	return e*2
})
fmt.Println(newArr) // => [2,4,6,8,10]
*/
func Map[T comparable](array []T, callback func(element T) T) []T {
	var newArr []T
	ForEach(array, func(e T) {
		newArr = append(newArr, callback(e))
	})
	return newArr
}

// Filter function for Array
/*
Usage:

arr := []int{1,2,3,4,5}
newArr := Filter(arr, func(e int) bool {
	return e%2==0
})
fmt.Println(newArr) // => [2,4]
*/
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
/*
Usage:

arr := []int{1,2,3,4,5}
sum := Reduce(arr, func(acc, e int) int {
	acc += e
	return acc
}, 0)
fmt.Println(sum) // => 15
*/
func Reduce[T, K comparable](array []T, callback func(accumulator K, element T) K, initial K) K {
	var acc = initial
	ForEach(array, func(e T) {
		acc = callback(acc, e)
	})
	return acc
}

// Reverse function for Array
/*
Usage:

arr := []int{1,2,3,4,5}
fmt.Println(Reverse(arr)) // => [5,4,3,2,1]
*/
func Reverse[T comparable](array []T) []T {
	var newArr []T
	for i := len(array) - 1; i >= 0; i-- {
		newArr = append(newArr, array[i])
	}
	return newArr
}

// Unique function for Array
/*
Usage:

arr := []int{1,2,2,2,3,4,4,4,4,5,5}
fmt.Println(Unique(arr)) // => [1,2,3,4,5]
*/
func Unique[T comparable](array []T) []T {
	var stagingArr []T
	ForEach(array, func(e T) {
		if Find(stagingArr, e) == -1 {
			stagingArr = append(stagingArr, e)
		}
	})
	return stagingArr
}

// Chunk function for array
/*
Usage: 

arr := []int{1,2,3,4,5}

fmt.Println(Chunk(arr, 2)) // => [[1 2] [3 4] [5]]
*/

func Chunk[T comparable](array []T, chunkSize int) [][]T{
	var newArr [][]T
	for i := 0; i < len(array); i += chunkSize{
		end := i + chunkSize
		if end > len(array){
			end = len(array)
		}
		newArr = append(newArr, array[i:end])
	}
	return newArr
}

// Flatten function for array of arrays
/*
Usage:

arr := [][]int{{1,2},{3,4},{5}}
fmt.Println(Flatten(arr)) // => [1 2 3 4 5]
*/
func Flatten[T comparable](array [][]T) []T{
	var newArr []T
	ForEach(array, func(sub []T) {
		ForEach(sub, func(e T) {
			newArr = append(newArr, e)
		})
	})
	return newArr
}