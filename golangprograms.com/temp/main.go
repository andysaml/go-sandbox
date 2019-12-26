package main

import (
	"fmt"
	"reflect"
)

type cat struct{}

func (c cat) say() {
	fmt.Println("meow")
}

type animal interface {
	say()
}

func main() {
	var a animal = cat{}
	fmt.Println(reflect.TypeOf(a)) // main.cat
	var b cat = a
	fmt.Println(reflect.TypeOf(b)) // main.cat
}
