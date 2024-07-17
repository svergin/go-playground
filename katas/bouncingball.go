package katas

// Kata: Bouncing Ball <6 kyu>
// https://www.codewars.com/kata/5544c7a5cb454edb3c000047
func CalcBounces(h, bounce, window float64) int {
	if h == 0 || (bounce == 0 || bounce >= 1) || window >= h {
		return -1
	}
	result := 1
	bh := h * bounce
	for {
		if bh < window {
			break
		}
		result = result + 2
		bh = bh * bounce
	}

	return result
}
