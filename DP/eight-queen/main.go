package main

import "fmt"

var count int = 0
var array *[8][8]int = new([8][8]int)

func Show() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if array[i][j] == 1 {
				fmt.Printf("x")
			} else if array[i][j] == 0 {
				fmt.Printf("o")
			}
		}
		fmt.Println()
	}
}

func Check(row, col int) bool {
	for i := 0; i <= row; i++ {
		for j := 0; j < 8; j++ {
			if (row != i || col != j) && array[i][j] == 1 {
				if row == i || col == j || row+col == i+j || row-col == i-j {
					return false
				}
			}
		}
	}

	return true
}

func Queen(row int) bool {

	if row > 7 {
		count++
		return true
	}

	for col := 0; col < 8; col++ {
		if Check(row, col) {

			if Queen(row + 1) {
				array[row][col] = 1
			}
		}
	}

	return false
}

func main() {

	Queen(0)

	//Show()

	fmt.Println(count)
}
