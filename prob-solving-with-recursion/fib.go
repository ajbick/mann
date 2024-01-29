package main

import (
	"fmt"
	"strconv"
)

var fibonacciValues []int64

func fibonacci(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciOnTheFly(n int64) int64 {
	if n == 0 {
		return fibonacciValues[n]
	}
	if n == 1 {
		return fibonacciValues[n]
	}
	if n >= int64(len(fibonacciValues)) {
		fibonacciValues = append(fibonacciValues, fibonacciOnTheFly(n-2)+fibonacciOnTheFly(n-1))
	}
	return fibonacciValues[n]
}

func main() {
	fibonacciValues = append(fibonacciValues, 0)
	fibonacciValues = append(fibonacciValues, 1)

	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)

		// Uncomment one of the following.
		fmt.Printf("fibonacciOnTheFly(%d) = %d\n", n, fibonacciOnTheFly(n))
		//fmt.Printf("fibonacciPrefilled(%d)  = %d\n", n, fibonacciPrefilled(n))
		//fmt.Printf("fibonacciBottomUp(%d)  = %d\n", n, fibonacciBottomUp(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacciValues); i++ {
		fmt.Printf("%d: %d\n", i, fibonacciValues[i])
	}
}
