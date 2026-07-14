package datastructures

type Set[T comparable] struct {
	registry map[T]any
	elements []T
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		registry: make(map[T]any),
		elements: make([]T, 0),
	}
}

func (s *Set[T]) Add(element T) {
	if _, exists := s.registry[element]; !exists {
		s.registry[element] = struct{}{}
		s.elements = append(s.elements, element)
		return
	}

	for idx, val := range s.elements {
		if val != element {
			continue
		}
		s.elements[idx] = element
		break
	}
}

func (s *Set[T]) Remove(element T) {
	if _, exists := s.registry[element]; !exists {
		return
	}

	delete(s.registry, element)
	for idx, val := range s.elements {
		if val != element {
			continue
		}
		s.elements = append(s.elements[:idx], s.elements[idx+1:]...)
		break
	}
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.registry[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.registry = make(map[T]any)
	s.elements = make([]T, 0)
}

func (s *Set[T]) ToSlice() []T {
	return s.elements
}
