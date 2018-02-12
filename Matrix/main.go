package main

import "fmt"

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

	array := make([][]int, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			array[i][j] = single(m1, m2, i, j)
		}
	}

	return array
}

func single(m1, m2 [][]int, x, y int) int {
	xlen := len(m1[0])
	ylen := len(m2)

	sum := 0

	for i := 0; i < xlen; i++ {
		for j := 0; j < ylen; j++ {

			sum += m1[x][i] * m2[j][y]
		}
	}

	return sum
}

func main() {

	fmt.Println("vim-go: ", multiplication(matrix1, matrix2))
}
