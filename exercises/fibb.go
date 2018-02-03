package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num1, num2 := 0, 1

	return func() int {
		result := num1
		// this doesnt work if separated--assignment at the same time, very cool :D
		num1, num2 = num2, num1 + num2
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

