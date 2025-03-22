package math

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func BenchmarkFibonacci(b *testing.B) {
	b.ReportAllocs()
	numCPU := runtime.NumCPU()
	fmt.Printf("Anzahl der CPU-Kerne: %d\n", numCPU)

	// Setze Go auf die maximale Anzahl von CPUs
	//runtime.GOMAXPROCS(numCPU)
	benchmarkFibo(b)

}
func BenchmarkFibonacciWithGoRoutines(b *testing.B) {
	numCPU := runtime.NumCPU()
	n := 40
	ch := make(chan time.Duration, numCPU)
	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go benchmarkFiboGoRoutine(n, &wg, ch)
	}
	wg.Wait()

	// Schließe den Channel, nachdem alle Goroutinen beendet sind
	close(ch)

	// Analysiere die Ergebnisse
	var totalDuration time.Duration
	for duration := range ch {
		totalDuration += duration
	}

	averageDuration := totalDuration / time.Duration(numCPU)
	fmt.Printf("Durchschnittliche Dauer für %d Benchmarks: %v\n", numCPU, averageDuration)
}

func benchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		fibonacci(40)
		d := time.Since(start)
		fmt.Printf("Dauer der Berechnung: %v\n", d)
	}
}

func benchmarkFiboGoRoutine(n int, wg *sync.WaitGroup, ch chan time.Duration) {

	defer wg.Done()
	start := time.Now()
	fibonacci(n)
	d := time.Since(start)
	ch <- d
}
