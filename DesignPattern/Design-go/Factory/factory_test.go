package factory

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("Design factory test!")
}

func Print(i IFather) {
	i.PrintName()
}

func TestFactory(t *testing.T) {
	var i IFatherFactory = new(AFatherFactory)
	var m IFatherFactory = new(BFatherFactory)

	a := i.CreateFather("panlin")
	Print(a)

	b := m.CreateFather("lijie")
	Print(b)
}
