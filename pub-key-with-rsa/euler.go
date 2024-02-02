package main

import (
	"fmt"
	//"time"
)

func largestOdd(x int) int {
	if x % 2 == 0 {
		return x - 1
	} else {
		return x
	}
}

func sieveOfEuler(max int) []bool {
	s := make([]bool, max+1)
	s[2] = true

	for i := 3; i < max; i += 2 {
		s[i] = true
	}

	for p := 3; p < max; p += 2 {
		if s[p] {
			maxQ := largestOdd(max / p)
			for q := maxQ; q >= p; q -= 2 {
				if s[q] {
					s[q*p] = false
				}
			}
		}
	}

	return s
}

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
	printSieve(sieveOfEratosthenes(100))
	printSieve(sieveOfEuler(100))
}
