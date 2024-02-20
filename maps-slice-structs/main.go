package main

import "fmt"

func main() {
	nomes := make([string, 10, 20]) // capacitando para 20 espaços no array
	idades := make(map[string]uint8)
	idades["Tiago"] = 31
	idades["Joana"] = 28
	idades["Maria"] = 35

	fmt.Println(idades)
}
