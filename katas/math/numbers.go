package math

import (
	"sync"
)

// Kata: Count the divisible numbers <6 kyu>
// https://www.codewars.com/kata/55a5c82cd8e9baa49000004c

type version string

const (
	V1 version = "v1"
	V2 version = "v2"
	V3 version = "v3"
)

func GetDividablesCount(v version, start, end, divider int) int {
	switch v {
	case V1:
		return getDividablesCount(start, end, divider)
	case V2:
		return getDividablesCountV2(start, end, divider)
	case V3:
		return getDividablesCountV3(start, end, divider)

	}
	return 0
}

func getDividablesCount(start, end, divider int) int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		if i%divider == 0 {
			result = append(result, i)
		}
	}
	return len(result)

}

func getDividablesCountV2(start, end, divider int) int {
	result := 0
	for i := start; i <= end; i++ {
		if i%divider == 0 {
			result++
		}
	}
	return result
}

func getDividablesCountV3(start, end, divider int) int {

	rntcnt := 8
	c := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= rntcnt; i++ {
		wg.Add(1)
		if i == 1 {
			go count(start, end/(rntcnt*i), divider, c, wg)
		} else {
			go count((end/rntcnt)*(i-1)+1, (end/rntcnt)*i, divider, c, wg)
		}
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	res := 0

	for i := range c {
		res = res + i
	}

	return res
}

func count(start, end, divider int, c chan int, wg *sync.WaitGroup) {
	result := 0
	for i := start; i <= end; i++ {
		if i%divider == 0 {
			result++
		}
	}
	c <- result
	wg.Done()

}
