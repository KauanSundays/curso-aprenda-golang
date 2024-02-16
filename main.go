package main

import "fmt"

var (
	nome string
	n1 int
	n2 int32
)

func somar(a int, b int) int{
	return a + b
}

func main() {
	var a, b, c = true, "toma", 123

	fmt.Println(a, b, c)
	fmt.Println(somar(2, 3))
}