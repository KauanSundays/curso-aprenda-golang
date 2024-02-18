package main

import "fmt"

func main() {
	nomes := []string{"Ana", "Bruno", "Carlos"}
	fmt.Println(nomes)

	nomes := append(nomes,  "Diana") //Adiciona um elemento no final da slice
	fmt.Println(nomes)
}
