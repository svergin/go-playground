package katas

type direction string

const North direction = "north"
const South direction = "south"
const West direction = "west"
const East direction = "east"

func DirReduc(plan []direction) []direction {
	for idx, cur := range plan {
		if len(plan) > idx+1 {
			next := plan[idx+1]
			if cur.isOposing(next) {
				plan = append(plan[:idx], plan[idx+2:]...)
				// DirReduc(plan)
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
