package list

import (
	"testing"
)

func BenchmarkListStringSven(b *testing.B) {
	var myList *List[string] = New("eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn")
	for i := 0; i < b.N; i++ {
		myList.String()
	}
}

func BenchmarkListStringAlex(b *testing.B) {
	var myList *List[string] = New("eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn")
	for i := 0; i < b.N; i++ {
		myList.String_alex()
	}
}
