package main

import "fmt"

func factorial(n int64) int64 {
	// enter your implementation there
	// calculation:
	// 1 for n <= 0
	// n! for n >= 1
	var product int64 = 1
	for i := n; i > 0; i-- {
		product = product * i
	}
	return product
}
func main() {
	fmt.Println(factorial(5))
}
