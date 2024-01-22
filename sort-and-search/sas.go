package main

import (
	"fmt"
	"math/rand"
	//"math"
	"time"
	"strconv"
)

type Customer struct {
    id           string
    numPurchases int
}

func makeRandomSlice(numItems, max int) []Customer {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]Customer, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = Customer{"C"+strconv.Itoa(i), rnd.Intn(max)}
	}
	return s
}

func makeRandomSliceInt(numItems, max int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rnd.Intn(max)
	}
	return s
}

func printSlice(slice []Customer, numItems int) {
	max := 0
	if len(slice) > numItems {
		max = numItems
	} else {
		max = len(slice)
	}
	fmt.Println(slice[:max])
}

func printSliceInt(slice []int, numItems int) {
	max := 0
	if len(slice) > numItems {
		max = numItems
	} else {
		max = len(slice)
	}
	fmt.Println(slice[:max])
}

func checkSorted(slice []Customer) {
	sz := len(slice)
	for i := 1; i < sz; i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted.")
}

func checkSortedInt(slice []int) {
	sz := len(slice)
	for i := 1; i < sz; i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return
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

func partition(slice []int) int {
	lo := 0
	hi := len(slice)-1
	
	pivot := slice[hi]

	i := lo - 1

	for j := lo; j < hi; j++ {
		if slice[j] <= pivot {
			i = i + 1
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]

	return i
}

func quicksort(slice []int) {
	if len(slice) < 2 {
		return
	}
	
	p := partition(slice)

	quicksort(slice[0:p])
	quicksort(slice[p+1:])
}

func countingSort(slice []Customer, max int) []Customer {
	counts := make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		counts[slice[i].numPurchases]++
	}

	for i := 1; i < len(slice); i++ {
		counts[i] += counts[i-1]
	}

	sorted := make([]Customer, len(slice))
	for i := len(slice)-1; i >= 0; i-- {
		x := slice[i]
		y := counts[x.numPurchases]
		sorted[y-1] = x
		counts[x.numPurchases]--
	}

	return sorted
}

func countingSortInt(slice []int, max int) []int {
	counts := make([]int, len(slice))

	for i := 0; i < len(slice); i++ {
		counts[slice[i]]++
	}

	for i := 1; i < len(slice); i++ {
		counts[i] += counts[i-1]
	}

	sorted := make([]int, len(slice))
	for i := len(slice)-1; i >= 0; i-- {
		x := slice[i]
		y := counts[x]
		sorted[y-1] = x
		counts[x]--
	}

	return sorted
}

func linearSearch(slice []int, target int) (index, numTests int) {
	for i, x := range slice {
		if x == target {
			return i, i
		}
	}
	return -1, len(slice)
}

func binarySearch(slice []int, target int) (index, numTests int) {
	l := 0
	r := len(slice)-1
	nt := 0

	for l <= r {
		m := (l + r) / 2
		nt++
		if slice[m] < target {
			l = m + 1
		} else if slice[m] > target {
			r = m - 1
		} else {
			return m, nt
		}
	}

	return -1, nt
}

func main() {
	var n, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&n)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	s := makeRandomSliceInt(n, max)
	//s := []int{0, 1, 1, 5, 6, 6, 6, 7, 10, 12, 14, 14, 15, 15, 17, 17, 20, 21, 21, 21}
	//fmt.Println(s)

	quicksort(s)
	printSliceInt(s, 40)	

	var tt string
	fmt.Printf("Target: ")
	fmt.Scanln(&tt)
	for tt != "q" {
		t, err := strconv.Atoi(tt)

		if err == nil {
			index, tests := binarySearch(s, t)
			if index > 0 {
				fmt.Println("slice[" + strconv.Itoa(index) + "] = " + strconv.Itoa(t) + ", " + strconv.Itoa(tests) + " tests.")
			} else {
				fmt.Println("Not found.")
			}
		}

		fmt.Printf("Target: ")
		fmt.Scanln(&tt)
	}
}
	
/*
// linear search
func main() {
	var n, max, t int
	fmt.Printf("# Items: ")
	fmt.Scanln(&n)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	s := makeRandomSliceInt(n, max)

	fmt.Println(s)

	fmt.Printf("Target: ")
	fmt.Scanln(&t)

	fmt.Println(linearSearch(s, t))
}

/*
// counting sort
func main() {
    // Get the number of items and maximum item value.
    var numItems, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&numItems)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted slice.
    slice := makeRandomSlice(numItems, max)
    printSlice(slice, 40)
    fmt.Println()

    // Sort and display the result.
    sorted := countingSort(slice, max)
    printSlice(sorted, 40)

    // Verify that it's sorted.
    checkSorted(sorted)
}

/*
// bubble sort
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
*/   

/*
// quick sort
func main() {
    // Get the number of items and maximum item value.
    var numItems, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&numItems)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted slice.
    slice := makeRandomSlice(numItems, max)
    printSlice(slice, 40)
    fmt.Println()

    // Sort and display the result.
    quicksort(slice)
    printSlice(slice, 40)

    // Verify that it's sorted.
    checkSorted(slice)
}
*/

