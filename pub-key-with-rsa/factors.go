package main

import (
	"fmt"
	"time"
)

var primes []int

func findFactors(num int) []int {
	var f []int

	for num % 2 == 0 {
		f = append(f, 2)
		num /= 2
	}

	factor := 3
	for factor * factor <= num {
		if num % factor == 0 {
			f = append(f, factor)
			num /= factor
		} else {
			factor += 2
		}
	}

	if factor != 1 {
		f = append(f, num)
	}

	return f
}

func findFactorsSieve(num int) []int {
	var f []int

	for _, factor := range primes {
		for num % factor == 0 {
			f = append(f, factor)
			num /= factor
			if num == 1 {
				break
			}
		}
		if factor * factor > num {
			f = append(f, num)
			break
		}
	}

	return f
}

func multiplySlice(s []int) int {
	result := 1
	for _, x := range s {
		result *= x
	}
	return s		
}

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

	primes = sieveToPrimes(sieveOfEuler(20000000))
	
	start1 := time.Now()
	fmt.Println(findFactors(312680865509917))
	elapsed1 := time.Since(start1)
	fmt.Printf("%f\n\n", elapsed1.Seconds())

	start2 := time.Now()
	fmt.Println(findFactorsSieve(312680865509917))
	elapsed2 := time.Since(start2)
	fmt.Printf("%f\n\n", elapsed2.Seconds())

}
