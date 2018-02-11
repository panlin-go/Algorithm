package factory

import "fmt"

func init() {
	fmt.Println("Design factory")
}

//Abstract class ------------------------------------------
type Father struct {
	Name string
}

func (f *Father) GetName() string {
	return f.Name
}

func (f *Father) PrintName() {
	fmt.Printf("Father name: %s\n", f.Name)
}

func (f *Father) SetName(name string) {
	f.Name = name
}

//factory AFather ------------------------------------------
type AFather struct {
	Father
}

//overwrite
func (f *AFather) PrintName() {
	fmt.Printf("AFather name: %s\n", f.Name)
}

//factory BFather ------------------------------------------
type BFather struct {
	Father
}

//overwrite
func (f *BFather) PrintName() {
	fmt.Printf("BFather name: %s\n", f.Name)
}

//
type AFatherFactory struct {
}

func (f *AFatherFactory) CreateFather(n string) IFather {
	return &AFather{Father{Name: n}}
}

//
type BFatherFactory struct {
}

func (f *BFatherFactory) CreateFather(n string) IFather {
	return &BFather{Father{Name: n}}
}

//Abstract interface ------------------------------------------
type IFatherFactory interface {
	CreateFather(name string) IFather
}

type IFather interface {
	PrintName()
}
