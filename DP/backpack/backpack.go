package main

import (
	"fmt"
	"math"
)

const MIN = 0.000001

var V [7]int = [7]int{35, 30, 60, 50, 40, 10, 25}
var W [7]int = [7]int{10, 40, 30, 52, 35, 40, 30}

/*
	d[num, total]
	d[3,10] = d[2,10],d[2,7]+12
	d[i,j] = max{d[i-1, j], d[i-1, j-v[i-1]]+W[i-1]}
*/
func Dynamic() int {

	D := new([7 + 1][150 + 1]int)

	for i := 0; i < 7+1; i++ {
		for j := 0; j < 150+1; j++ {
			if i == 0 {
				D[i][j] = 0
			} else {
				D[i][j] = D[i-1][j]
			}

			if i > 0 && j > V[i-1] && D[i-1][j] < D[i-1][j-V[i-1]]+W[i-1] {
				D[i][j] = D[i-1][j-V[i-1]] + W[i-1]
			}
		}
	}

	return D[7][150]
}

func Sort(array *[7]float64, len int) {

	for i := 0; i < len; i++ {
		for j := 0; j < len-1-i; j++ {
			if math.Dim(array[j], array[j+1]) < MIN {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func Greed() int {
	S := new([7]float64)

	for i := 0; i < 7; i++ {
		S[i] = float64(W[i]) / float64(V[i])
	}

	Sort(S, 7)

	fmt.Println(S)

	return 0
}

func main() {
	Greed()

	fmt.Println("result: ", Dynamic())
}
