package ccl

type Set[T comparable] struct {
	_items []T
}

func CreateSet[T comparable]() Set[T] {
	return Set[T]{}
}

func CreateSetFromArray[T comparable](array []T) Set[T] {
	SetArr := Unique(array)
	return Set[T]{
		_items: SetArr,
	}
}

func (s *Set[T]) Add(element T) {
	if Find(s._items, element) == -1 {
		s._items = append(s._items, element)
	}
}

func (s Set[T]) Array() []T {
	return s._items
}
