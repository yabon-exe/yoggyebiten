package model

type Size[T int | uint | float32 | float64] struct {
	W T
	H T
}

func (s *Size[T]) Set(w T, h T) {
	s.W = w
	s.H = h
}

func (s *Size[T]) Swap() {
	temp := s.W
	s.W = s.H
	s.H = temp
}

func (s *Size[T]) ToWH() (T, T) {
	return s.W, s.H
}
