package slist

import (
	"errors"
	"fmt"
)

type Data int

type Listnode struct {
	data Data
	next *Listnode
}

func NewList() *Listnode {
	return &Listnode{next: nil}
}

func (head *Listnode) Insert(mode string, data Data) error {
	if head == nil {
		return errors.New("head is nil!")
	}

	node := &Listnode{data: data, next: nil}
	var list *Listnode = head

	switch mode {
	case "tail":
		{
			for list.next != nil {
				list = list.next
			}

			list.next = node
		}

	case "head":
		{
			node.next = list.next
			list.next = node
		}

	default:
		{
			return errors.New("mode must head|tail!")
		}
	}

	return nil
}

func (head *Listnode) InsertList(mode string, data ...Data) error {
	num := len(data)

	for i := 0; i < num; i++ {
		if err := head.Insert(mode, data[i]); err != nil {
			return err
		}
	}

	return nil
}

func (head *Listnode) Delete(mode string, data Data) error {
	if head == nil {
		return errors.New("head is nil!")
	}

	var flag bool = false
	var list *Listnode = head

	for list.next != nil {
		if list.next.data == data {
			flag = true
			list.next = list.next.next

			switch mode {
			case "single":
				{
					break
				}
			case "more":
				{
					//
				}
			default:
				{
					return errors.New("mode must single|more!")
				}
			}

		} else {
			list = list.next
		}
	}

	if flag != true {
		return errors.New("data not found!")
	}

	return nil
}

func (head *Listnode) DeleteList(mode string, data ...Data) error {
	num := len(data)
	for i := 0; i < num; i++ {
		if err := head.Delete(mode, data[i]); err != nil {
			return err
		}
	}

	return nil
}

func (head *Listnode) ShowList() error {
	if head == nil {
		return errors.New("head is nil!")
	}

	var list *Listnode = head

	fmt.Printf("head")
	list = list.next

	for list != nil {
		fmt.Printf("->%d", list.data)
		list = list.next
	}

	fmt.Println()
	return nil
}

/*func main() {

	head := NewList()

	head.InsertList("head", 10, 5, 2, 3, 5, 7)

	head.ShowList()
	head.DeleteList("more", 5, 3)

	head.ShowList()
}*/
