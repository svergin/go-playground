package katas

import (
	"errors"
	"strconv"
	"strings"
)

// Kata: Count IP Addresses <5 kyu>
// https://www.codewars.com/kata/526989a41034285187000de4

type ip struct {
	first  int64
	second int64
	third  int64
	fourth int64
}

func CountIPs(startIP, endIP string) (int64, error) {
	var s, e ip
	var err error
	if s, err = convert(startIP); err != nil {
		return 0, err
	}
	if e, err = convert(endIP); err != nil {
		return 0, err
	}
	if s.first > e.first {
		return 0, errors.New("Startwert größer als Endwert")
	}

	fir := val(s.first, e.first)
	sec := val(s.second, e.second)
	thi := val(s.third, e.third)
	fou := val(s.fourth, e.fourth)

	return fir + sec + thi + fou, nil
}

func val(a, b int64) int64 {
	if a == 0 && b > 0 {
		return b
	}

	if a != b {
		return 255 - a + b
	}

	return 0
}

func convert(ipStr string) (ip, error) {
	var r = ip{}
	var err error
	parts := strings.Split(ipStr, ".")
	if r.first, err = strconv.ParseInt(parts[0], 10, 32); err != nil {
		return r, err
	}
	if r.second, err = strconv.ParseInt(parts[1], 10, 32); err != nil {
		return r, err
	}
	if r.third, err = strconv.ParseInt(parts[2], 10, 32); err != nil {
		return r, err
	}
	if r.fourth, err = strconv.ParseInt(parts[3], 10, 32); err != nil {
		return r, err
	}
	return r, nil

}
