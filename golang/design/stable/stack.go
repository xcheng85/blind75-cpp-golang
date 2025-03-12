package stable

// interface for stack generic

// c++ template
// template<typename T>
// class stack

type Stack[T any] interface {
	Top() T
	Empty() bool
	Size() int
	Push(T)
	Pop()
}
