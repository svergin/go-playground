package math

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Kata: Factorial decomposition
// https://www.codewars.com/kata/5a045fee46d843effa000070
func Decomp(n int) string {

	num := calcFalcuty(n)

	primes := getPrimes(num)
	res := make(map[int]int, len(primes))
	count := 0
	tmpnum := num

	for i := 0; i < len(primes); i++ {
		for {
			if tmpnum%primes[i] != 0 {
				break
			}
			tmpnum = tmpnum / primes[i]
			count++
		}
		if count > 0 {
			res[primes[i]] = count
			count = 0
		}
	}

	return toResultString(res)
}

func DecompV2(facultyNum int) string {

	facultyResult := calcFalcuty(facultyNum)

	primes := getPrimesV2(facultyResult)

	res := make(map[int]int, len(primes))
	count := 0
	tmpnum := facultyResult

	for i := 0; i < len(primes); i++ {
		for {
			if tmpnum%primes[i] != 0 {
				break
			}
			tmpnum = tmpnum / primes[i]
			count++
		}
		if count > 0 {
			res[primes[i]] = count
			count = 0
		}
	}

	return toResultString(res)
}

func DecompGoRout(facultyNum int) string {

	facultyResult := calcFalcuty(facultyNum)

	primes := getPrimes(facultyResult)

	result := make(map[int]int, len(primes))
	countPrimeChan := make(chan primecount, 1)

	for i := 0; i < len(primes); i++ {
		go countPrimeFactors(countPrimeChan, facultyResult, primes[i])
		p := <-countPrimeChan
		if p.prime > 0 {
			result[p.prime] = p.count
		}

	}

	return toResultString(result)
}

func toResultString(primeFactorMapping map[int]int) string {
	keys := make([]int, 0, len(primeFactorMapping))

	for k := range primeFactorMapping {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sb := strings.Builder{}
	cnt := 0
	for _, k := range keys {
		v := primeFactorMapping[k]

		if v == 1 {
			sb.WriteString(fmt.Sprintf("%v", k))
		} else {
			sb.WriteString(fmt.Sprintf("%v^%v", k, v))
		}
		cnt++
		if cnt < len(keys) {
			sb.WriteString(" * ")
		}
	}

	return sb.String()
}

type primecount struct {
	prime int
	count int
}

func countPrimeFactors(primeChan chan primecount, rootNum int, prime int) {
	tmpnum := rootNum
	result := primecount{}
	for {
		if tmpnum%prime != 0 {
			break
		}
		tmpnum = tmpnum / prime
		result.count++
	}
	if result.count > 0 {
		result.prime = prime
	}
	primeChan <- result
}

func calcFalcuty(n int) int {
	if n == 0 {
		return 1
	}
	num := 1
	for i := n; i >= 1; i-- {
		num = num * i
	}
	return num
}

func getPrimesV2(num int) []int {
	// start := time.Now()
	length := num + 1
	boolMatrix := make([]bool, length)

	result := make([]int, 0)
	// Sieb des Erastothenes v2
	lim := int(math.Sqrt(float64(length)))
	for i := 2; i <= lim; i++ {
		if !boolMatrix[i] {
			result = append(result, i)
			for j := i * i; j < length; j = j + i {
				boolMatrix[j] = true
			}
		}
	}

	for i := lim + 1; i < length; i++ {
		if !boolMatrix[i] {
			result = append(result, i)
		}

	}
	// fmt.Printf("getPrimesV2() dauerte %v ms\n", time.Since(start).Milliseconds())
	return result
}

func getPrimes(num int) []int {
	// start := time.Now()
	if num < 2 {
		return []int{}
	}
	if num == 2 {
		return []int{2}
	}
	knownPrimes := []int{2, 3, 5, 7, 11}
	// start := time.Now()
	numToFilter := make([]int, num-4)
	// fmt.Printf("Erstellung Zahlenreihen-Slice dauerte %v ms\n", time.Since(start).Milliseconds())

	// start = time.Now()
	for i := range numToFilter {
		numToFilter[i] = 3 + i
	}
	// fmt.Printf("Befüllung Zahlenreihen-Slice dauerte %v ms\n", time.Since(start).Milliseconds())

	var relevantPrimes []int

	for _, p := range knownPrimes {
		if p <= num {
			relevantPrimes = append(relevantPrimes, p)
		}
	}
	// start = time.Now()
	result := filterPrimes(numToFilter, relevantPrimes)
	// fmt.Printf("Filterung der Zahlenreihen-Slice dauerte %v ms\n", time.Since(start).Milliseconds())
	// fmt.Printf("getPrimes() dauerte %v ms\n", time.Since(start).Milliseconds())

	return result
}

// Sieb des Erastothenes v1
func filterPrimes(numToFilter, knownPrimes []int) []int {

	resMap := make(map[int]int, len(knownPrimes))
	for _, p := range knownPrimes {
		resMap[p] = p
	}
	for _, n := range numToFilter {
		if _, ok := resMap[n]; ok {
			continue
		}

		isPrime := true
		for _, p := range knownPrimes {

			if p > n {
				isPrime = false
				break
			}
			if n%p == 0 && n != p {
				isPrime = false
				break
			}
		}
		if isPrime {
			resMap[n] = n
		}
	}
	result := make([]int, 0, len(resMap))
	for k := range resMap {
		result = append(result, k)
	}
	return result
}
