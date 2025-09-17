package ccl

type Set[T comparable] struct {
	_items []T
}

func CreateSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Creates Set from existing array
func CreateSetFromArray[T comparable](array []T) Set[T] {
	setArr := Unique(array)
	return Set[T]{
		_items: setArr,
	}
}

func (s *Set[T]) Add(element T) {
	if Find(s._items, element) == -1 {
		s._items = append(s._items, element)
	}
}

// Converts the Set to Array
func (s Set[T]) Array() []T {
	return s._items
}
