package main

import "fmt"

func main() {
	nomes := []string{"tom Brady", "Lamar Jackson", "Joe Burrow"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}
}