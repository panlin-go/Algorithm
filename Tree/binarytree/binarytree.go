package main

import (
	"errors"
	"fmt"
)

type Data int

type Tree struct {
	data  Data
	right *Tree
	left  *Tree
	heigh int
}

func NewTree(data Data) *Tree {
	return &Tree{data: data, right: nil, left: nil}
}

func (t *Tree) InsertRight(data Data) error {
	if t == nil {
		return errors.New("tree is nil!")
	}

	node := &Tree{data: data, right: nil, left: nil}

	t.right = node

	return nil
}

func (t *Tree) InsertLeft(data Data) error {
	if t == nil {
		return errors.New("tree is nil!")
	}

	node := &Tree{data: data, right: nil, left: nil}

	t.left = node

	return nil
}

func (t *Tree) Insert(data Data) error {
	if t == nil {
		return errors.New("tree is nil!")
	}

	for i := t; i != nil; {
		if data > i.data {
			if i.right == nil {
				node := &Tree{data: data, right: nil, left: nil}
				i.right = node
				break
			} else {
				i = i.right
			}
		} else {
			if i.left == nil {
				node := &Tree{data: data, right: nil, left: nil}
				i.left = node
				break
			} else {
				i = i.left
			}
		}
	}

	return nil
}

func (t *Tree) InsertList(data ...Data) error {
	for _, v := range data {
		if err := t.Insert(v); err != nil {
			return err
		}
	}
	return nil
}

func Traversal(t *Tree) {
	if t == nil {
		return
	}

	fmt.Printf("%d\t\n", t.data)

	Traversal(t.right)
	Traversal(t.left)
}

func main() {
	tree := NewTree(5)

	if err := tree.InsertList(2, 4, 6, 8); err != nil {
		fmt.Println("insert error")
	}

	Traversal(tree)
}
