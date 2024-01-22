package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeRandomSlice(numItems, max int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rnd.Intn(max)
	}
	return s
}

func printSlice(slice []int, numItems int) {
	max := 0
	if len(slice) > numItems {
		max = numItems
	} else {
		max = len(slice)
	}
	fmt.Println(slice[:max])
}

func checkSorted(slice []int) {
	sz := len(slice)
	for i := 1; i < sz; i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
		}
	}
	fmt.Println("The slice is sorted.")
}

func bubbleSort(slice []int) {
	for iter := 0; iter < len(slice); iter++ {
		for i := 1; i < len(slice)-iter; i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
			}
		}
	}
}

func main() {
    // Get the number of items and maximum item value.
    var numItems, max int
    fmt.Printf("# Items: ")
    fmt.Scanln(&numItems)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display an unsorted slice.
    slice := makeRandomSlice(numItems, max)
    printSlice(slice, 40)
    fmt.Println()

    // Sort and display the result.
    bubbleSort(slice)
    printSlice(slice, 40)

    // Verify that it's sorted.
    checkSorted(slice)
}
