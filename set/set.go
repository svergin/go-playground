package set

type Set[T comparable] map[T]T // Welcher Typ soll verwendet werden

// Ein Element hinzufügen mit O(1)
func (s *Set[T]) Add(v ...T) {
	var mySet Set[T]
	if s == nil {
		mySet = make(map[T]T)
	} else {
		mySet = *s
	}

	for _, wert := range v[0:] {
		if mySet.Contains(wert) {
			continue
		}
		mySet[wert] = wert
	}
	*s = mySet
}

// Prüfen, ob ein Element vorhanden ist mit O(1)
func (s *Set[T]) Contains(v T) bool {
	if s == nil {
		return false
	}
	var mySet Set[T] = *s
	_, exists := mySet[v]
	return exists

}

func (s *Set[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}
