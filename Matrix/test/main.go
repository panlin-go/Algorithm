package main

import (
	"fmt"
	"reflect"
)

func main() {

	row := 3
	//col := 3

	temp := make([]int, row)
	ttype := reflect.TypeOf(temp)
	fmt.Println("type:", ttype)
	//array := make([]ttype, col)

	//fmt.Println(array)
}
