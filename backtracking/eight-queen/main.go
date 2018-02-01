package main

import (
	"fmt"
	"os"
	"strconv"
)

var ROW int = 0
var COL int = 0

var Count int = 0

func Check(array []int, row, col int) bool {
	for i := 0; i < row; i++ {
		x := array[i]
		y := i

		//   |            ——               /             \
		if row == y || col == x || row-col == y-x || row+col == y+x {
			return false
		}
	}

	return true
}

func Show(array []int) {
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

func Queen(array []int, row int) {
	if row > (ROW - 1) {
		Count++

		// show map
		Show(array)
		return
	}

	for i := 0; i < COL; i++ {
		//check node is vaild
		if Check(array, row, i) {
			array[row] = i
			Queen(array, row+1)
		}
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Uasge: %s num\n", os.Args[0])
		return
	} else {
		ROW, _ = strconv.Atoi(os.Args[1])

		COL, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println("begin count: ", Count)

	array := make([]int, ROW)

	Queen(array, 0)

	fmt.Println("after count: ", Count)
}
