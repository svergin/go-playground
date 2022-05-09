package list

import (
	"fmt"
	"strings"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) Value() T {
	return l.val
}

func (l List[T]) Tail() *List[T] {
	return l.next
}

func New[T any](vals ...T) *List[T] {
	if len(vals) == 0 {
		return nil
	}

	return &List[T]{
		val:  vals[0],
		next: New(vals[1:]...),
	}

}

func New_sv[T any](vals ...T) *List[T] {

	fmt.Println("Called with: ", vals, len(vals), cap(vals))
	var l List[T]
	for i, wert := range vals {
		// fmt.Printf("Vals: %v :: i: %d\n", vals, i)
		l = List[T]{val: wert}
		i++
		if len(vals) > 1 {
			l.next = New_sv(vals[i:]...)
		} else {
			l.next = nil
		}
		return &l
	}
	return &l
}

func (l List[T]) String() string {
	var sb strings.Builder
	var naechster *List[T] = l.Tail()
	var wert any = l.Value()
	for {
		result := fmt.Sprintf("Wert: %v, naechster: %p --> ", wert, naechster)
		sb.WriteString(result)
		if naechster != nil {
			wert = naechster.Value()
			naechster = naechster.Tail()
		} else {
			break
		}

	}
	return sb.String()
}

func (l *List[T]) String_alex() string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%v -- %s", l.val, l.next.String())
}
