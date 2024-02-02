package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p int, numTests int) bool {
	
}

func randRange(min int, max int) int {
	return min + random.Intn(max - min)
}

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

func main() {

}
