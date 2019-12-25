package main

import "fmt"

type Cat struct {
}

type Dog struct {
}

type Sound interface {
	say()
}

func (cat Cat) say() {
	fmt.Println("meoy!")
}

func (dog Dog) say() {
	fmt.Println("bark!")
}

func main() {
	a := Cat{}
	b := Dog{}
	arr := []Sound{a, b}

	for _, item := range arr {
		item.say()
	}
}
