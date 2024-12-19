package collections

type LNode[T any] struct {
	next  *LNode[T]
	value T
}

func NewLNode[T any](v T) *LNode[T] {
	return &LNode[T]{next: nil, value: v}
}

func (l *LNode[T]) InsertAfter(v T) {
	l.next = &LNode[T]{next: l.next, value: v}
}

func (l *LNode[T]) Value() T {
	return l.value
}

func (l *LNode[T]) SetValue(v T) {
	l.value = v
}

func (l *LNode[T]) Next() *LNode[T] {
	return l.next
}

func (l *LNode[T]) SetNext() *LNode[T] {
	return l.next
}
