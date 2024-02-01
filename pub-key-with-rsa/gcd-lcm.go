package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	if a < 0 {
		a = a * -1
	}
	if b < 0 {
		b = b * -1
	}
	
	if b > a {
		a, b = b, a
	}

	if a == b {
		return b
	} else {
		return gcd(a-b, b)
	}
}

func lcm(a int, b int) int {
	if a < 0 {
		a = a * -1
	}
	if b < 0 {
		b = b * -1
	}

	return (a * b) / gcd(a, b)
}


func main() {
	fmt.Println(gcd(270, 192))
	fmt.Println(lcm(270, 192))
	fmt.Println()
	fmt.Println(gcd(7469, 2464))
	fmt.Println(lcm(7469, 2464))
	fmt.Println()
	fmt.Println(gcd(55290, 115430))
	fmt.Println(lcm(55290, 115430))
	fmt.Println()
	
}
