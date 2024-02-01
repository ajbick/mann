package main

import (
	"fmt"
)

func fastExp(num int, pow int) int {
	result := 1
	for pow > 0 {
		if pow % 2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

func fastExpMod(num int, pow int, mod int) int {
	result := 1
	for pow > 0 {
		if pow % 2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

func main() {
	for {
		var num int
		var pow int
		var mod int

		fmt.Printf("Num:")
		fmt.Scan(&num)
		if num < 1 {
			break
		}

		fmt.Printf("Pow:")
		fmt.Scan(&pow)
		if pow < 1 {
			break
		}

		fmt.Printf("Mod:")
		fmt.Scan(&mod)
		if mod < 1 {
			break
		}

		fmt.Printf("fastExp    = %d\n", fastExp(num, pow))
		fmt.Printf("fastExpMod = %d\n", fastExpMod(num, pow, mod))
		fmt.Println()
	}
}

	
	
