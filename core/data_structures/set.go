package datastructures

type Set[T comparable] struct {
	elements map[T]any
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]any),
	}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for element := range s.elements {
		slice = append(slice, element)
	}
	return slice
}
