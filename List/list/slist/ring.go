package slist

import (
	"errors"
	"fmt"
)

func (head *Listnode) CreateRing(data Data) error {
	if head == nil {
		return errors.New("head is nil!")
	}

	var flag bool = false

	var p *Listnode = head
	var ringnode *Listnode
	var tailnode *Listnode

	//
	for p.next != nil {
		if p.next.data == data {
			flag = true
			ringnode = p.next
		}

		p = p.next
	}

	if flag == false {
		return errors.New("data not found!")
	} else {
		tailnode = p
	}

	tailnode.next = ringnode

	return nil
}

func (head *Listnode) DetectRing() (bool, error) {
	if head == nil {
		return false, errors.New("head is nil!")
	}

	var fast *Listnode = head.next
	var slow *Listnode = head.next

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			count := 1
			slow = head.next

			for {
				if fast == slow {
					break
				}

				fast = fast.next
				slow = slow.next
				count++
			}

			fmt.Println("ring data: ", fast.data, "\tcount: ", count)

			return true, nil
		}
	}

	return false, nil
}
