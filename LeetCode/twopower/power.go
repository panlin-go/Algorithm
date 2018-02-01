package main

import (
	"fmt"
	"os"
	"strconv"
)

func TwoP(n int) bool {
	return n&(n-1) == 0
}

func main() {
	len := len(os.Args)
	if len != 2 {
		fmt.Println("Usage: ./coin 11")
		return
	}

	num, _ := strconv.Atoi(os.Args[1])

	fmt.Println("result: ", TwoP(num))
}
