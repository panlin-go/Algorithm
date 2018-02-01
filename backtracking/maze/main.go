package main

import (
	"container/list"
	"fmt"
)

var array [6][6]int = [6][6]int{
	{1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 0, 0},
	{1, 0, 0, 1, 1, 1},
	{1, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1}}

var Count int = 0

type Point struct {
	row int
	col int
}

func ShowList(l *list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		if f, ok := e.Value.(Point); ok {
			fmt.Printf("[%d,%d]->", f.row, f.col)
		}
	}

	fmt.Println()
}

func Check(l *list.List, target Point) bool {
	element := l.Front().Next()

	if element == nil {
		return true
	}

	if f, ok := element.Value.(Point); ok {
		if target.row == f.row && target.col == f.col {
			return false
		}
		return true
	}

	return false
}

func Find(src, dst Point, l *list.List) {
	if src.row > 6-1 || src.col > 6-1 || src.row < 0 || src.col < 0 || array[src.row][src.col] != 1 || array[dst.row][dst.col] != 1 {
		return
	}

	// add node list
	l.PushFront(src)

	if src.row == dst.row && src.col == dst.col {
		Count++
		ShowList(l)

		// add node
		l.Remove(l.Front())
		return
	}

	//up
	if Check(l, Point{src.row - 1, src.col}) {
		Find(Point{src.row - 1, src.col}, dst, l)
	}
	//down
	if Check(l, Point{src.row + 1, src.col}) {
		Find(Point{src.row + 1, src.col}, dst, l)
	}
	//right
	if Check(l, Point{src.row, src.col + 1}) {
		Find(Point{src.row, src.col + 1}, dst, l)
	}
	//left
	if Check(l, Point{src.row, src.col - 1}) {
		Find(Point{src.row, src.col - 1}, dst, l)
	}

	// add node
	l.Remove(l.Front())
}

func main() {
	l := list.New()

	Find(Point{0, 0}, Point{5, 5}, l)

	fmt.Println("count: ", Count)
}
