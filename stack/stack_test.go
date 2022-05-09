package stack

import (
	"fmt"
	"testing"
)

func TestStackPeek(t *testing.T) {

	testStack := New('S', 'V', 'E', 'N')
	got := testStack.Peek()
	expected := 'N'
	if *got != expected {
		t.Errorf("Stack.Peek() = %v; want %v", *got, expected)
	}
}

func TestStackPushViaValue(t *testing.T) {
	testStack := New("T", "E", "S")
	testStack = testStack.Push_Value("T")
	got := testStack.Peek()
	expected := "T"
	if *got != expected {
		t.Errorf("Stack.Push() = %v; want %v", *got, expected)
	}
}

func TestStackPushViaPointer(t *testing.T) {
	testStack := New("T", "E", "S")
	testStack.Push("T")
	fmt.Println("testStack nach PushP: ", testStack)
	got := testStack.Peek()
	expected := "T"
	if *got != expected {
		t.Errorf("Stack.Push() = %v; want %v", *got, expected)
	}
}

// func TestStackPop(t *testing.T) {
// 	testStack := New("T", "E", "S")
// 	got := testStack.Pop()
// 	fmt.Println("Nach POP: ", testStack)
// 	expected := "S"
// 	if *got != expected {
// 		t.Errorf("Stack.Pop() = %v; want %v", *got, expected)
// 	}
// 	gotPeek := testStack.Peek()
// 	expectedPeek := "E"
// 	if *gotPeek != expectedPeek {
// 		t.Errorf("Stack.Peek() = %v; want %v", *gotPeek, expectedPeek)
// 	}
// }

func TestStackPopViaPointer(t *testing.T) {
	testStack := New("T", "E", "S")
	got, err := testStack.Pop()
	if err != nil {
		t.Error(err)
	}
	expected := "S"
	if *got != expected {
		t.Errorf("Stack.Pop() = %v; want %v", *got, expected)
	}
	gotPeek := testStack.Peek()
	expectedPeek := "E"
	if *gotPeek != expectedPeek {
		t.Errorf("Stack.Peek() = %v; want %v", *gotPeek, expectedPeek)
	}
}

func TestStackPopOnEmptyStack(t *testing.T) {
	testStack := New(1)
	_, err := testStack.Pop()
	_, err = testStack.Pop()

	if err == nil {
		t.Error("Error expected, Stack was empty")
	}
}
