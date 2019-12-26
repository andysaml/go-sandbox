package main

import (
	"fmt"
	"reflect"
)

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	fmt.Println(reflect.TypeOf(p))         // main.Pentagon
	fmt.Println(reflect.ValueOf(p).Kind()) // int
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	fmt.Println(reflect.TypeOf(o))         // main.Pentagon
	fmt.Println(reflect.ValueOf(o).Kind()) // int
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	fmt.Println(reflect.TypeOf(obj))         // main.Pentagon
	fmt.Println(reflect.ValueOf(obj).Kind()) // int
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	fmt.Println(reflect.TypeOf(pent))         // main.Pentagon
	fmt.Println(reflect.ValueOf(pent).Kind()) // int
	pent.Perimeter()
}
