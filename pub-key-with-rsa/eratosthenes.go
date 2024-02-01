package main

import (
	"fmt"
	"time"
)

func sieveOfEratosthenes(max int) []bool {
	a := make([]bool, max+1)

	a[2] = true

	for i := 3; i < max; i += 2 {
		a[i] = true
	}

	for i := 3; i < max; i += 2 {
		if a[i] {
			for j := i * 3; j <= max; j += i {
				a[j] = false
			}
		}
	}

	return a
}

func printSieve(sieve []bool) {
	fmt.Printf("2 ")
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func sieveToPrimes(sieve []bool) []int {
	var s []int
	s = append(s, 2)
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			s = append(s, i)
		}
	}
	return s
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
