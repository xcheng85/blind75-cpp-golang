package unstable

import "errors"

type SliceStackAdaptor[T any] struct {
	data []T
}


func (s *SliceStackAdaptor[T]) Top() (T, error) {
	if !s.Empty() {
		return s.data[0], nil
	} else {
		var noop T
		return noop, errors.New("stack is empty")
	}
}

func (s *SliceStackAdaptor[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *SliceStackAdaptor[T]) Size() int {
	return len(s.data)
}

func (s *SliceStackAdaptor[T]) Push(v T) {
	s.data = append([]T{v}, s.data...)
}

func (s *SliceStackAdaptor[T]) Pop() (res T, err error) {
	if!s.Empty() {
		res = s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	var noop T
	return noop, errors.New("stack is empty")
}

func NewSliceStackAdaptor[T any]() *SliceStackAdaptor[T] {
	return &SliceStackAdaptor[T]{
		data: make([]T, 0),
	}
}
