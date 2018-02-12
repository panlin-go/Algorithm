package main

import "fmt"

func minarray_enum(array []int) ([]int, int) {

	var start int = 0
	var stop int = 0

	var max int = 0

	for i := 0; i < len(array); i++ {
		sum := 0
		for j := i; j < len(array); j++ {
			sum += array[j]

			if sum > max {
				start = i
				stop = j

				max = sum
			}
		}
	}

	fmt.Println(start, stop)
	return array[start : stop+1], max
}

func middiemax(array []int, left, rigth, middle int) int {
	lmax, rmax := 0, 0

	sum := 0
	for i := middle + 1; i < rigth; i++ {
		sum += array[i]
		if sum > rmax {
			rmax = sum
		}
	}

	sum = 0
	for j := middle; j > left; j-- {
		sum += array[j]
		if sum > lmax {
			lmax = sum
		}
	}

	return lmax + rmax
}

func divide(array []int, left, right int) int {
	sum := array[left]

	if left < right {
		middle := (right - left) / 2
		lmax, rmax, mmax := 0, 0, 0

		lmax = divide(array, left, middle)
		rmax = divide(array, middle, left)
		mmax = middiemax(array, left, right, middle)
		sum = mmax

		if lmax > rmax && lmax > mmax {
			sum = lmax
		}

		if rmax > lmax && rmax > mmax {
			sum = rmax
		}
	}

	return sum
}

func minarray_divide(array []int) int {

	return divide(array, 0, len(array)-1)
}

/*
	d[0] = 0
	d[1] = array[1-1]
	d[2] = max {d[2-1], d[2-1]+array[2-1], array[2-1]}
	d[n] = {d[n-1] > 0 ? max{d[n-1], d[n-1] + array[2-1]} : max{array[n-1]}}
*/
func minarray_dp(array []int) ([]int, int) {
	start, stop := 0, 0

	sum := 0
	d := make([]int, len(array)+1)

	d[0] = 0

	for i := 1; i < len(d); i++ {
		d[i] = d[i-1]
		if sum < 0 {
			if array[i-1] > d[i-1] {
				d[i] = array[i-1]
				sum = d[i]
				start, stop = i-1, i-1
			}
		} else {
			sum += array[i-1]
			if sum > d[i-1] {
				stop = i - 1
				d[i] = sum
			}
		}
	}

	return array[start : stop+1], d[len(d)-1]
}

func main() {

	array := []int{-2, 5, 3, -6, 4, -8, 6}

	//marr, result := minarray_enum(array)
	marr, result := minarray_dp(array)

	result1 := minarray_divide(array)

	fmt.Println("array: ", marr)
	fmt.Println("result: ", result)
	fmt.Println("result1: ", result1)
}
