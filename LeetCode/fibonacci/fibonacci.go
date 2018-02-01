package main

/*
	f(x) = f(x-1)+f(x-2)  f(1) = 1  f(2) = 1


*/

import (
	"fmt"
	"os"
	"strconv"
)

func Recursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return Recursion(n-1) + Recursion(n-2)
}

func Dynamic(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		} else {
			return 1
		}
	}

	array := make([]int, n+1)

	array[1] = 1
	array[2] = 1

	for i := 3; i < n+1; i++ {
		array[i] = array[i-1] + array[i-2]
	}

	return array[n]
}

func main() {
	len := len(os.Args)
	if len != 2 {
		fmt.Println("Usage: ./coin 11")
		return
	}

	num, _ := strconv.Atoi(os.Args[1])

	result := Dynamic(num)
	fmt.Println("Dynamic result: ", result)

	result = Recursion(num)
	fmt.Println("Recursion result: ", result)
}
