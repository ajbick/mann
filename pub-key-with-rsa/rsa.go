package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

func randRange(min int, max int) int {
	return min + random.Intn(max-min)
}

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

func totient(p int, q int) int {
	return lcm(p-1, q-1)
}

func randomExponent(n int) int {
	for {
		e := randRange(2, n)
		if gcd(e, n) == 1 {
			return e
		}
	}
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

func inverseMod(a int, n int) int {
	t := 0
	r := n
	newt := 1
	newr := a

	for newr != 0 {
		q := r / newr
		t, newt = newt, t-q*newt
		r, newr = newr, r-q*newr
	}

	if r > 1 {
		return -1
	}
	if t < 0 {
		t = t + n
	}

	return t
}

func main() {
	// Pick two random primes p and q.
	const numTests = 100
	p := findPrime(10000, 50000, numTests)
	q := findPrime(10000, 50000, numTests)

	// Calculate the public key modulus n.
	n := p * q

	// Calculate Carmichael's totient function
	tn := totient(p, q)

	// Pick a random public key exponent e in the range [3, tn)
	// where gcd(e, tn) = 1.
	e := randomExponent(tn)

	// Find the inverse of e mod tn.
	d := inverseMod(e, tn)

	// Print out the important values.
	fmt.Printf("*** Public ***\n")
	fmt.Printf("Public key modulus:    %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)
	fmt.Printf("\n*** Private ***\n")
	fmt.Printf("Primes:    %d, %d\n", p, q)
	fmt.Printf("t(n):      %d\n", tn)
	fmt.Printf("d:         %d\n", d)

	for {
		var m int
		fmt.Printf("\nMessage:    ")
		fmt.Scan(&m)
		if m < 1 {
			break
		}

		ciphertext := fastExpMod(m, e, n)
		fmt.Printf("Ciphertext: %d\n", ciphertext)

		plaintext := fastExpMod(ciphertext, d, n)
		fmt.Printf("Plaintext:  %d\n", plaintext)
	}
}
