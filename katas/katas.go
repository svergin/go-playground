package katas

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Kata: Beginner - Lost Without a Map
func Maps(x []int) []int {
	output := make([]int, len(x))

	for pos, v := range x {
		output[pos] = v * 2
	}
	return output
}

// Kata: Square(n) Sum
func SquareSum(numbers []int) int {
	var result int = 0

	for _, v := range numbers {
		result += v * v
	}
	return result
}

// Kata: Remove String Spaces
func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

// Kata: Sum of positive
func PositiveSum(numbers []int) int {
	var result int = 0
	for _, v := range numbers {
		if v >= 0 {
			result += v
		}
	}
	return result
}

// Kata: Heron's formula
func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return area
}

// Kata: Get the Middle Character
func GetMiddle(s string) string {

	laenge := len(s)
	if laenge <= 1 {
		return s
	}
	stringslice := strings.Split(s, "")
	if laenge%2 == 1 {
		return stringslice[(laenge-1)/2]
	} else {
		haelfte := laenge / 2
		return strings.Join(stringslice[haelfte-1:haelfte+1], "")
	}
}

// Kata: Vowel Count
func GetCount(str string) (count int) {
	stringslice := strings.Split(str, "")

	for _, v := range stringslice {
		switch v {
		case "a", "i", "e", "o", "u":
			count++
		}

	}

	return count
}

// Kata: Highest and Lowest
func HighAndLow(in string) string {
	hlSlice := strings.Split(in, " ")
	var min int64 = 1
	var max int64 = 1
	for _, v := range hlSlice {
		zahl, _ := strconv.ParseInt(v, 10, 32)
		if zahl < min {
			min = zahl
		}
		if zahl > max {
			max = zahl
		}
	}
	return fmt.Sprint(max, " ", min)
}

// Kata: Mumbling
func Accum(s string) string {
	builder := strings.Builder{}
	toAccum := strings.Split(s, "")
	var repeater int = 0
	for count, v := range toAccum {
		builder.WriteString(strings.ToUpper(v))
		builder.WriteString(strings.Repeat(strings.ToLower(v), repeater))
		repeater++
		if count < len(toAccum)-1 {
			builder.WriteString("-")
		}
	}

	return builder.String()
}

// Kata: Shortest Word
func FindShortest(str string) int {
	var minLength = len(str)
	words := strings.Split(str, " ")
	for _, v := range words {
		var rpl = strings.NewReplacer(".", "", "?", "", "!", "", ";", "")
		v = rpl.Replace(v)
		if len(v) < minLength {
			minLength = len(v)
		}
	}
	return minLength
}

// Kata: Two to One
func TwoToOne(s1, s2 string) string {
	s1s := strings.Split(s1, "")
	allLetters := append(s1s, strings.Split(s2, "")...)

	letterMap := make(map[string]string)
	for _, v := range allLetters {
		letterMap[v] = v
	}
	uniqueLetters := make([]string, 0, len(letterMap))
	for k := range letterMap {
		uniqueLetters = append(uniqueLetters, k)
	}
	sort.Strings(uniqueLetters)
	builder := strings.Builder{}
	for _, k := range uniqueLetters {
		builder.WriteString(k)
	}

	return builder.String()
}

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
	sb := strings.Builder{}
	// fmt.Sprintln(pnums)
	cnt := 0
	for k, v := range res {
		sb.WriteString(fmt.Sprintf("%v^%v", k, v))
		cnt++
		if cnt < len(res) {
			sb.WriteString(" * ")
		}
	}
	return sb.String()
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
	sb := strings.Builder{}
	cnt := 0
	for k, v := range result {
		sb.WriteString(fmt.Sprintf("%v^%v", k, v))
		cnt++
		if cnt < len(result) {
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

//primeCache := map[int]int {2:2}

func getPrimes(num int) []int {
	if num < 2 {
		return []int{}
	}
	if num == 2 {
		return []int{2}
	}
	knownPrimes := []int{2, 3, 5, 7, 11}
	numToFilter := make([]int, num-4)

	for i := range numToFilter {
		numToFilter[i] = 3 + i
	}

	var relevantPrimes []int

	for _, p := range knownPrimes {
		if p <= num {
			relevantPrimes = append(relevantPrimes, p)
		}
	}

	result := filterPrimes(numToFilter, relevantPrimes)

	return result
}

// Sieb des Erastothenes
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
