package main

import "unicode/utf8"

func main() {
	a := "abc"
	println(a)
	println(len(a))

	a = "абц"
	println(a)
	println(len(a))
	println(utf8.RuneCountInString(a))

}
