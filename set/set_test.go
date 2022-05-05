package set

import "testing"

func TestUninitializedSetLengthIsZero(t *testing.T) {
	//Schlägt nicht fehl, wenn NIL-Prüfung in len() fehlt. Warum?!
	var mySet Set[string]
	var got int = mySet.Len()
	var expected int = 0
	if got != expected {
		t.Errorf("Got: %v, length should be %v", got, expected)
	}
}
func TestSetLengthIsZero(t *testing.T) {
	var mySet Set[string] = make(Set[string], 5)
	var got int = mySet.Len()
	var expected int = 0
	if got != expected {
		t.Errorf("Got: %v, length should be %v", got, expected)
	}
}

func TestSetLengthGreaterZero(t *testing.T) {
	var mySet Set[string] = make(Set[string], 5)
	mySet.Add("eins")
	var got int = mySet.Len()
	var expected int = 1
	if got != expected {
		t.Errorf("Got: %v, length should be %v", got, expected)
	}
}

func TestSetContainsElement(t *testing.T) {
	var mySet Set[string] = make(Set[string], 5)
	mySet.Add("eins", "zwei", "drei")
	var got bool = mySet.Contains("eins")
	if !got {
		t.Errorf("Got: %v, element should be contained in set", got)
	}

}

func TestSetDoesntContainElement(t *testing.T) {
	var mySet Set[string] = make(Set[string], 5)
	mySet.Add("eins", "zwei", "drei")
	var got bool = mySet.Contains("vier")
	if got {
		t.Errorf("Got: %v, element should not be contained in set", got)
	}

}
