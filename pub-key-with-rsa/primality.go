package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	return min + random.Intn(max-min)
}

func fastExpMod(num int, pow int, mod int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

func findPrime(min int, max int, numTests int) int {
	for {
		p := randRange(min, max)
		if p%2 == 0 {
			continue
		}

		if isProbablyPrime(p, numTests) {
			return p
		}
	}
}

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p int, numTests int) bool {
	for i := 0; i < numTests; i++ {
		n := randRange(1, p)

		x := fastExpMod(n, p-1, p)

		if x != 1 {
			return false
		}
	}
	return true
}

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
var numTests = 20

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf("Probability: %f%%\n\n", (1.0-1.0/math.Pow(2, float64(numTests)))*100.0)

	fmt.Println("Primes:")
	for _, number := range primes {
		if isProbablyPrime(number, 10) {
			fmt.Println(number, " Prime")
		} else {
			fmt.Println(number, " Composite")
		}
	}
	fmt.Println()

	fmt.Println("Composites:")
	for _, number := range composites {
		if isProbablyPrime(number, 10) {
			fmt.Println(number, " Prime")
		} else {
			fmt.Println(number, " Composite")
		}
	}
}

func main() {
	//testKnownValues()
	fmt.Println(findPrime(1, 100, 20))
}
