package main

import "fmt"

const (
	i         = 10
	f         = 1.5
	i64 int64 = 88
)

const (
	Black = iota
	Gray
	White
)

// счётчик обнуляется
const (
	Yellow = iota
	Red
	Green = iota // это присваивание не обнулит iota
	Blue
)

func main() {
	ver := "v0.0.1"
	id := 0
	pi := 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	var v = 45
	println(i + f)
	println(i + i64)
	//i3 := f + i64
	//i4 := i64 + v
	println(i + v)
	//i6 := f + v

	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue)
}
