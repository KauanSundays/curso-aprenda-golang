package main

import "fmt"

type Pessoa struct {
	Nome 		string
	Sobrenome  	string
	Idade      	int
	Status     	bool
}

type Categoria struct {
	Nome 		string
	Pai *Categoria
}

// func (c Categoria) HasParent() bool {
// 	return c.Pai != nil 
// }

func main() {
	p := Pessoa{
		Nome:		"Tiago",
		Sobrenome:	"Alves",
		Idade: 		30,
		Status: 	true,
	}

	fmt.Println(p) // imprime o conteúdo da estrutura de dados
}