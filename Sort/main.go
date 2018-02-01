package main

import "fmt"

func show(array []int) {
	n := len(array)

	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", array[i])
	}

	fmt.Println("")
}

func maopao(array []int) {
	n := len(array)

	flag := true

	for i := 1; i < n && flag; i++ {
		flag = false

		for j := 0; j < n-i; j++ {
			flag = true
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	show(array)
}

func kuaipai(array []int, low, high int) {
	temp := array[low]
	i := high
	j := low

	for i > j {
		for ; i > low && i > j; i-- {
			if array[i] < temp {
				array[j] = array[i]
				j++
				break
			}
		}

		for ; j < high && i > j; j++ {
			if array[j] > temp {
				array[i] = array[j]
				i--
				break
			}
		}
	}

	array[j] = temp

	if low < j-1 {
		kuaipai(array, low, j-1)
	}
	if j+1 < high {
		kuaipai(array, j+1, high)
	}

	show(array)
}

func main() {
	array := []int{38, 5, 19, 26, 49, 97, 1, 66}

	//maopao(array)
	kuaipai(array, 0, 7)
}
