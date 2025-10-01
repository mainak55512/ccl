# Common Collection Library – Generic Data Structures & Functional Utilities for Go

A lightweight collection of **functional utilities** and **data structures** implemented in **Go Generics**.  
This project provides handy operations for arrays, enums, and sets, making it easier to work with collections in a functional programming style.  

---

## ✨ Features  

- **Array Utilities**  
  - `ForEach` → Iterate over elements  
  - `Find` → Find an element in an array  
  - `Map` → Transform an array  
  - `Filter` → Filter elements based on condition  
  - `Reduce` → Accumulate values  
  - `Reverse` → Reverse an array  
  - `Unique` → Remove duplicates  
  - `Chunk` → Split array into chunks of given size
  - `Flatten` → Flatten 2D array into 1D array

- **Enum Implementation**  
  - Create enums with auto-incremented or custom values  
  - Freeze enums to prevent further modification  
  - Pattern matching with exhaustive checks  

- **Set Data Structure**  
  - Add unique elements  
  - Create sets from arrays  
  - Convert sets back to arrays  
  - Merge two Sets with Union method
  - Get common elements of two Sets with Intersection method
  - Get difference of two Sets with Difference method

---

## 📦 Installation  

```bash
go get github.com/mainak55512/ccl@latest
```

Then import into your project:  

```go
import "github.com/mainak55512/ccl"
```

---

## 🚀 Usage  

### 🔹 Array Utilities  

```go
arr := []int{1, 2, 3, 4, 5}

// ForEach
ccl.ForEach(arr, func(e int) {
	fmt.Println(e)
})

// Find
fmt.Println(ccl.Find(arr, 4)) // => 3 (index) or -1 if not found

// Map
doubled := ccl.Map(arr, func(e int) int { return e * 2 })
fmt.Println(doubled) // => [2,4,6,8,10]

// Filter
evens := ccl.Filter(arr, func(e int) bool { return e%2 == 0 })
fmt.Println(evens) // => [2,4]

// Reduce
sum := ccl.Reduce(arr, func(acc, e int) int { return acc + e }, 0)
fmt.Println(sum) // => 15

// Reverse
fmt.Println(ccl.Reverse(arr)) // => [5,4,3,2,1]

// Unique
dupArr := []int{1,2,2,3,4,4,5}
fmt.Println(ccl.Unique(dupArr)) // => [1,2,3,4,5]

// Chunk
fmt.Println(ccl.Chunk(arr, 2)) // => [[1 2] [3 4] [5]]

// Flatten
nested := [][]int{{1, 2}, {3, 4}, {5}}
fmt.Println(ccl.Flatten(nested)) // => [1,2,3,4,5]
```

---

### 🔹 Enum  

```go
en := ccl.CreateEnum()
en.AddWithValue("SUCCESS", 5)
en.Add("FAILURE")
en.Add("PENDING")
en.Add("IN_PROGRESS")
en.Freeze()

status := 6

state, err := en.Match(status, map[string]func() any{
	"SUCCESS":     func() any { return 1 },
	"FAILURE":     func() any { return 2 },
	"PENDING":     func() any { return 3 },
	"IN_PROGRESS": func() any { return 4 },
	"_DEFAULT_":   func() any { return 0 },
})

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("State:", state) // State: 2
}
```

---

### 🔹 Set  

```go
// Create Sets from Array
s1 := ccl.CreateSetFromArray([]string{"India", "Australia", "Brazil", "America"})
s2 := ccl.CreateSetFromArray([]string{"Russia", "Brazil", "America", "China", "Endland", "Portugal"})

// Union method returns Set of all unique elements from both sets
fmt.Println(s1.Union(s2).Array())        // => [China Endland Australia Brazil America India Portugal Russia]

// Intersection returns Set of common elements from both sets
fmt.Println(s1.Intersection(s2).Array()) // => [Brazil America]

// Difference returns Set of elements of current Set that are not present in other set
fmt.Println(s1.Difference(s2).Array())   // => [Australia India]
fmt.Println(s2.Difference(s1).Array())   // => [China Endland Portugal Russia]
```

---

## 📖 Roadmap  

- [ ] Add unit tests  
- [ ] Add more array helpers (Chunk, Flatten, etc.)  
- [ ] Extend Enum with string-based support  
- [ ] Optimize Set with map-based storage  

---

## 🤝 Contributing  

Contributions, issues, and feature requests are welcome!  
Feel free to open a [PR or an issue](./CONTRIBUTING.md).  

---

## 📜 License  

MIT License © 2025 [Mainak Bhattacharjee]  
