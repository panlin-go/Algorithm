package sort

import "fmt"

func init() {
	fmt.Println("Func list:")
	fmt.Println("MergeSort")
}

func Merge(array []int, sub1 []int, sub2 []int) []int {
	if len(sub1) == 1 && len(sub2) == 1 {

		arr := make([]int, 2)

		if sub1[0] > sub2[0] {
			arr[0] = sub2[0]
			arr[1] = sub1[0]
		} else {
			arr[0] = sub1[0]
			arr[1] = sub2[0]
		}
		fmt.Println(array)
		fmt.Println(sub1)
		fmt.Println(sub2)

		fmt.Println(arr)
		return arr
	}

	fmt.Println("=============================")
	fmt.Println(array)
	fmt.Println(sub1)
	fmt.Println(sub2)

	array1 := Merge(sub1, sub1[0:len(sub1)/2], sub1[len(sub1)/2:len(sub1)])
	array2 := Merge(sub2, sub2[0:len(sub2)/2], sub2[len(sub2)/2:len(sub2)])

	//merge
	var i1, i2 int = 0, 0
	a := make([]int, len(sub2)+len(sub1))

	for j := 0; j < len(a); j++ {
		if len(array1) > i1 && len(array2) > i2 {
			if array1[i1] > array2[i2] {
				a[j] = array2[i2]
				i2++
			} else {
				a[j] = array1[i1]
				i1++
			}
		} else if len(array1) > i1 {
			a[j] = array1[i1]
			i1++
		} else if len(array2) > i2 {
			a[j] = array2[i2]
			i2++
		} else {
			fmt.Println("error!")
		}
	}
	fmt.Println(a)

	return a
}

func MergeSort(array []int) {
	result := Merge(array, array[0:len(array)/2], array[len(array)/2:len(array)])
	for i := 0; i < len(result); i++ {
		array[i] = result[i]
	}
}
