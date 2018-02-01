package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	l.PushFront(6)
	l.PushFront(5)
	l.PushFront(4)
	l.PushFront(3)
	l.PushFront(2)

	for e := l.Front(); e != nil; e = e.Next() {
		if f, ok := e.Value.(int); ok {
			fmt.Println(f)
		}
	}

	fmt.Println("vim-go")
}
