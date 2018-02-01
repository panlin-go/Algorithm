package main

/*
	xor: n^n = 0 /  n^n^n = n
*/

import "fmt"

var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5}

func Xor() int {
	temp := 0

	for i := 0; i < 10; i++ {
		temp = temp ^ array[i]
	}

	for i := 1; i < 10; i++ {
		temp = temp ^ i
	}

	return temp
}

func main() {

	fmt.Println("result: ", Xor())

}
