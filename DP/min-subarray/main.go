package main

import "fmt"

/*
	V{10, -17,70...}
	d[0] = 0
	d[1] = V[0]
	d[2] = max{d[1],d[1]+V[1]}

	d[n] = max{d[n-1], d[n-1]+V[n-1]}
*/
func Dynamic(array []int) []int {
	len := len(array)

	d := make([]int, len+1)

	for i := 0; i < len+1; i++ {
		d[i] = i
		if i > 0 {
			if d[i-1] > d[i-1]+array[i-1] {
				d[i] = d[i-1]
			} else {
				d[i] = array[i-1] + d[i-1]
			}
		}
	}
	return d
}

func Transition(array []int) []int {
	len := len(array)

	result := make([]int, len-1)

	for i := 0; i < len-1; i++ {
		result[i] = array[i+1] - array[i]
	}

	return result
}

func ShowArray(array []int) {
	for _, v := range array {
		fmt.Printf("%d\t", v)
	}

	fmt.Println()
}

func main() {
	var array []int = []int{10, 22, 5, 75, 65, 80}

	r := Transition(array)
	ShowArray(r)

	d := Dynamic(r)

	fmt.Println(d[5])
}
