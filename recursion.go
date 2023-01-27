package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fib(number-1) + fib(number-2)
	return result
}
