package sort

import "fmt"

func init() {
	fmt.Println("Func list:")
	fmt.Println("InsertSort")
}

/*
	5,2,7,9,1,3
	[5],2,7,9,1,3
	[2,5],7,9,1,3
*/
func InsertionSort(array []int) {
	len := len(array)
	var i, j int

	for i = 1; i < len; i++ {
		//select key
		temp := array[i]
		for j = i - 1; j >= 0 && temp < array[j]; j-- {
			//[2,6,8],5  -> [2,6,8],8
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}
