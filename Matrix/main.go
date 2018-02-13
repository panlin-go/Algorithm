package main

import (
	"fmt"
)

var matrix1 = [][]int{
	{1, 2},
	{1, -1}}

var matrix2 = [][]int{
	{1, 2, -3},
	{-1, 1, 2}}

/*
	C[i,j] = A-row[i] * B-col[j]
*/
func multiplication(m1, m2 [][]int) [][]int {
	row := len(m1)
	col := len(m2[0])

	array := make([][]int, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum := 0
			for m := 0; m < row; m++ {
				sum += m1[i][m] * m2[m][j]
			}
			array[i] = append(array[i], single(m1, m2, i, j))
		}
	}

	return array
}

func main() {

	fmt.Println("vim-go: ", multiplication(matrix1, matrix2))
}
