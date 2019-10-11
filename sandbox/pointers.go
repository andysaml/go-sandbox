package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}



/*
В Go есть указатели. Указатель содержит адрес переменной в памяти.
Тип *T это указатель(адрес переменной в памяти) на значение T. Его нулевое значение nil.
var p *int
Оператор & возвращает указатель на его операнд.
i := 42
p = &i
Оператор * дает доступ к нижележащему значению указателя.
fmt.Println(*p) // прочитать i через указатель p
*p = 21         // установить i через указатель p
Это известно как "разыменование" и "присваивание".
В отличие от C, в Go нет адресной арифметики.
*/