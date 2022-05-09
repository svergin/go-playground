package stack

import (
	"fmt"
)

type Stack[T any] []T

//Erstellt einen Stack (LIFO)
func New[T any](vals ...T) *Stack[T] {
	var myStack Stack[T] = vals[0:]
	return &myStack
}

func (s Stack[T]) PushOri(t T) {
	//TODO
}

// Ein neues Element auf den Stack legen
func (s Stack[T]) Push_Value(t T) *Stack[T] {
	newLength := len(s) + 1
	if newLength > cap(s) {
		s = append(s, t)
	} else {
		s = s[:newLength]
		s[newLength] = t
	}
	return &s
}

// Ein neues Element auf den Stack legen
func (s *Stack[T]) Push(t T) {
	var s1 = *s
	newLength := len(s1) + 1
	if newLength > cap(s1) {
		s1 = append(s1, t)
	} else {
		s1 = s1[:newLength]
		s1[newLength] = t
	}
	*s = s1
}

/* Funktioniert nicht, da das Element aus dem kopierten Stack
entfernt wird und der Original-Stack unverändert bleibt
*/
func (s Stack[T]) Pop_Bad() *T {
	last := &s[len(s)-1]
	s = s[0 : len(s)-1]
	fmt.Println("POP: ", s)
	return last
}

// Das neueste Element aus dem Stack entfernen (LIFO)
//TODO: Wie funtioniert die Fehlerbehandlung, wenn T kein Pointer ist?
func (s *Stack[T]) Pop() (*T, error) {

	if len(*s) == 0 {
		fmt.Println(len(*s) == 0)

		return nil, fmt.Errorf("Unable to execute Pop(), Stack is empty")
	}
	s1 := *s
	last := &s1[len(s1)-1]
	*s = s1[0 : len(s1)-1]
	fmt.Println("POPP: ", s)
	return last, nil
}

func (s *Stack[T]) PopPVreturn() T {

	s1 := *s
	last := s1[len(s1)-1]
	*s = s1[0 : len(s1)-1]
	fmt.Println("POPP: ", s)
	return last
}

// Prüfen, ob ein Element aktuell auf dem Stack liegt
func (s Stack[T]) Peek() *T {
	return &s[len(s)-1]
}
