package katas

type direction string

const North direction = "north"
const South direction = "south"
const West direction = "west"
const East direction = "east"

func DirReduc(plan []direction) []direction {

	for i := 0; i < len(plan); i++ {
		if len(plan) > i+1 {
			cur := plan[i]
			next := plan[i+1]
			if cur.isOposing(next) {
				plan = append(plan[:i], plan[i+2:]...)
				i = -1
			}
		}
	}

	return plan
}

func (d direction) isOposing(other direction) bool {
	if (d == North && other == South) || (d == South && other == North) {
		return true
	}
	if (d == West && other == East) || (d == East && other == West) {
		return true
	}
	return false
}
