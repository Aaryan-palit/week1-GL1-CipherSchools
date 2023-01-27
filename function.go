package main

import "fmt"

/*
func main() {
	statement-1
	statement-2
	statement-3
	statement-4
}
*/

func main() {
	result, x := c()
	fmt.Println(result)
	fmt.Println(result)
	fmt.Println(x)
	r, _ := b(1, 2, 3, 4, 5, 6, 6)
	fmt.Println(r)
}

func a() (int, string) {
	return 122, "weqrr3"
}

func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c() (i int, j string) {
	i = 10
	j = "rahul"
	return
}
