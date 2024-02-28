package stack

type Stack[T any] struct {
	keys []T
	tail int
}

func New[T any](size int) *Stack[T] {
	return &Stack[T]{keys: make([]T, size), tail: 0}
}

func (stack *Stack[T]) Push(key T) {
	stack.tail++
	stack.keys[stack.tail] = key
}

func (stack *Stack[T]) Pop() (T, bool) {
	var x T
	if stack.tail == 0 {
		return x, false
	}

	return x, true
}
