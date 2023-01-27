package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	nums := []int{1, 3, 2, 4, 0}
	for index, value := range nums {
		fmt.Println(index)
		fmt.Println(value)
	}
	for _, value := range "rahul" {
		// if value== "u"{
		// 	break;
		// }
		fmt.Println(value)
	}
}
