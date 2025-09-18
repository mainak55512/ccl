package ccl

type Set[T comparable] struct {
	_items map[T]struct{}
}

// Creates an empty set
func CreateSet[T comparable]() Set[T] {
	return Set[T]{
		_items: make(map[T]struct{}),
	}
}

// Converts an array to Set
func CreateSetFromArray[T comparable](array []T) Set[T] {
	set := CreateSet[T]()
	ForEach(array, func(element T) {
		set.Add(element)
	})
	return set
}

// Adds element to set
func (s *Set[T]) Add(element T) {
	s._items[element] = struct{}{}
}

// Converts the Set to an Array
func (s Set[T]) Array() []T {
	items := make([]T, 0, len(s._items))
	for k := range s._items {
		items = append(items, k)
	}
	return items
}

// Merges two sets
func (s Set[T]) Union(compareSet Set[T]) Set[T] {
	newSet := CreateSet[T]()
	for k := range s._items {
		newSet._items[k] = struct{}{}
	}
	for k := range compareSet._items {
		newSet._items[k] = struct{}{}
	}
	return newSet
}

// Returns a Set with common elements
func (s Set[T]) Intersection(compareSet Set[T]) Set[T] {
	newSet := CreateSet[T]()
	for k := range s._items {
		for m := range compareSet._items {
			if k == m {
				newSet.Add(k)
				break
			}
		}
	}
	return newSet
}

// Returns a sub-set of the current set where the elements are not present in the other set
func (s Set[T]) Difference(compareSet Set[T]) Set[T] {
	newSet := CreateSet[T]()
	for k := range s._items {
		found := false
		for m := range compareSet._items {
			if k == m {
				found = true
				break
			}
		}
		if !found {
			newSet.Add(k)
		}
	}
	return newSet
}
