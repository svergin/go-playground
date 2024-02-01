package katas

import "fmt"

// Kata: Human Readable Time <5 kyu>
// https://www.codewars.com/kata/52685f7382004e774f0001f7

type hrtime struct {
	hours   uint
	minutes uint
	seconds uint
}

const MaxValue uint = 359999

func toHumanReadableTime(seconds uint) (string, error) {
	if seconds > MaxValue {
		return "", fmt.Errorf("%d seconds exceed maximum value of %d", seconds, MaxValue)
	}
	res := hrtime{}
	if seconds >= 60 {
		min := seconds / 60
		res.seconds = seconds % 60
		res.minutes = min % 60
		res.hours = min / 60
		if min < 60 {
			res.minutes = min
		}
	} else {
		res.seconds = seconds
	}
	return res.String(), nil
}

func (t hrtime) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.hours, t.minutes, t.seconds)
}
