package main

import "fmt"

func main() {
	var a int = 5
	p := &a

	fmt.Println(a, p)

	i := 42
	p1 := &i
	fmt.Println(*p1) // читаем значение переменной i через указатель p
	*p1 = 21         // записываем в переменную i значение 21 через указатель p
	fmt.Println(i)   // читаем значение переменной i
	fmt.Println(*p1) // читаем значение переменной i через указатель p

	a1 := 1
	p2 := &a1
	b := &p2

	*p2 = 3
	**b = 5

	a1 += 4 + *p2 + **b

	fmt.Printf("%d", *p2)

}
