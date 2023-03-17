package main

import "fmt"

type A struct {
	IntField int
}

func main() {
	p := &A{
		IntField: 10,
	}
	fmt.Println(p)
}
