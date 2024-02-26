package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ { // inicia a variavel, vezes, continuidade
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond + 150)
	}
	done <- true
}

func letras() {
	for l := 'a'; l < 'j'; l++ { // inicia a variavel, vezes, continuidade
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond + 170)
	}
	done <- true
}

func main() {
	cn := make(chan bool)
	cl := make(chan bool)
	go numeros(cn) // routines
	go letras(cl)

	<-cn
	<-cl

	time.Sleep(3 * time.Second) // Espera
}
