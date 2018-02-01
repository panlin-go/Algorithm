package main

import (
	"arithmetic/list/slist"
	_ "fmt"
)

func main() {
	head := slist.NewList()

	head.InsertList("head", 10, 5, 2, 3, 5, 7)

	head.ShowList()

	head.Reversal()

	head.ShowList()

	//	head.CreateRing(2)
	//
	//	b, err := head.DetectRing()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	if b == true {
	//		fmt.Println("Detect list have ring!")
	//	}
	//
	head.ReversalShow()
	return
}
