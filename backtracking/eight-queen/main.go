package main

import (
	"fmt"
	_ "os"
)

const (
	ROW int = 8
	COL int = 8
)

var Count int = 0

func Check(array *[ROW]int, row, col int) bool {
	for i := 0; i < row; i++ {
		x := array[i]
		y := i

		if row == y || col == x || row-col == y-x || row+col == y+x {
			return false
		}
	}

	return true
}

func Show(array *[ROW]int) {
	fmt.Printf("count: %d ===================\n", Count)

	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if j == array[i] {
				fmt.Printf("%2.c", 'X')
			} else {
				fmt.Printf("%2.c", 'O')
			}
		}

		fmt.Println()
	}
}

func Queen(array *[ROW]int, row int) {
	if row > (ROW - 1) {
		Count++

		// show map
		Show(array)
		return
	}

	for i := 0; i < COL; i++ {
		if Check(array, row, i) {
			array[row] = i
			Queen(array, row+1)
		}
	}

}

func main() {

	fmt.Println("begin count: ", Count)

	array := new([ROW]int)

	Queen(array, 0)

	fmt.Println("after count: ", Count)
}
