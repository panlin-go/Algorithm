package adjacency

import (
	"fmt"
)

type List struct {
	Name string
	Next *List
}

type Member struct {
	Count int
	Name  string
	Next  *List
}

type Adjacency struct {
	Count   int
	Member  []Member
	InputDU map[string]int
}

func New() *Adjacency {
	ad := new(Adjacency)
	ad.Count = 0
	ad.Member = make([]Member, 0)
	ad.InputDU = make(map[string]int, 0)
	return ad
}

func (ad *Adjacency) CheckRepetition(name string) bool {
	for i := 0; i < len(ad.Member); i++ {
		if ad.Member[i].Name == name {
			return true
		}
	}

	return false
}

func (ad *Adjacency) Add(name string, input ...string) bool {
	if ok := ad.CheckRepetition(name); ok {
		fmt.Println("Add failed. result: name is exist")
		return false
	}

	mem := new(Member)
	mem.Name = name
	mem.Count = 0
	mem.Next = nil

	if _, ok := ad.InputDU[mem.Name]; !ok {
		ad.InputDU[mem.Name] = 0
	}
	//
	for i := 0; i < len(input); i++ {
		if input[i] == mem.Name {
			fmt.Println("Add failed. result: input == name")
			return false
		}

		list := new(List)
		list.Name = input[i]

		tmp := mem.Next
		mem.Next = list
		list.Next = tmp

		mem.Count++
		//
		if _, ok := ad.InputDU[list.Name]; ok {
			ad.InputDU[list.Name]++
		} else {
			ad.InputDU[list.Name] = 1
		}
	}

	ad.Member = append(ad.Member, *mem)
	ad.Count++

	return true
}

func CheckExist(mem Member, name string) bool {
	if mem.Name == name {
		return false
	}

	var p *List = mem.Next
	for i := 0; i < mem.Count; i++ {
		if p.Name == name {
			return true
		}

		p = p.Next
	}

	return false
}

/*
func (ad *Adjacency) GenerateInputDU() bool {
	for i := 0; i < len(ad.Member); i++ {
		input := 0
		name := ad.Member[i].Name
		for j := 0; j < len(ad.Member); j++ {
			if ok := CheckExist(ad.Member[j], name); ok {
				input++
			}
		}
		ad.InputDU[name] = input
	}

	return true
}
*/

func (ad *Adjacency) Show() {
	for i := 0; i < len(ad.Member); i++ {
		mem := ad.Member[i]
		fmt.Printf("[%s]: ", mem.Name)
		var p *List = mem.Next
		for j := 0; j < mem.Count; j++ {
			fmt.Printf("%s ", p.Name)
			p = p.Next
		}
		fmt.Println()
	}

	fmt.Println("InputDU:")
	for k, v := range ad.InputDU {
		fmt.Printf("[%s]=%d\n", k, v)
	}
}
