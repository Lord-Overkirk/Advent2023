package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func fac(n int) (res int) {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}

func sqrt(x float64) (z float64) {
	z = 1.0
	n_guesses := 10
	for i := 0; i < n_guesses; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println((z))
	}
	return
}

func main() {
	var a int = add(42, 13)
	fmt.Println(a)
	fmt.Println(fac(10))
	sqrt(25)
}
