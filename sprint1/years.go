package main

import "fmt"

func check(year int) {
	switch {
	case year < 1964:
		fmt.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Привет, представитель X!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year >= 1997 && year < 2012:
		fmt.Println("Привет, зумер!")
	default:
		fmt.Println("Привет, альфа!")
	}
}
func main() {
	check(1946)
	check(1965)
	check(1981)
	check(1997)
	check(2013)
}
