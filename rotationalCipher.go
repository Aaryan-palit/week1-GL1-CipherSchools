package main

import "fmt"

var alphabet = map[int]string{
	1:  "a",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "r",
	19: "s",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",
}

func GetKey(letter string) (key int) {
	for index, value := range alphabet {
		if value == letter {
			key = index
			break
		}
	}
	return
}

//shifting it five times for beginner

func RotationCipher(original string) (encryptic string) {
	for _, value := range original {
		shiftedkey := GetKey(string(value)) + 5
		if shiftedkey > 26 {
			shiftedkey = shiftedkey % 26
		}
		encryptic += alphabet[shiftedkey]
	}
	return
}

func main() {
	fmt.Println(RotationCipher("omg"))
}
