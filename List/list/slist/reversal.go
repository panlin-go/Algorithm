package slist

import (
	"errors"
	"fmt"
)

func (head *Listnode) Reversal() error {
	if head == nil {
		return errors.New("head is nil!")
	}

	var p1 *Listnode = head.next
	var p2 *Listnode = p1.next
	var p3 *Listnode

	p1.next = nil

	for p2 != nil {
		p3 = p2.next

		p2.next = p1

		p1 = p2

		p2 = p3
	}

	head.next = p1
	return nil
}

//recursion have stack
func (head *Listnode) ReversalShow() error {
	if head == nil {
		return errors.New("head is nil!")
	}

	temp := head.next

	if temp == nil {
		return nil
	}

	temp.ReversalShow()

	fmt.Printf("%d->", temp.data)
	return nil
}
