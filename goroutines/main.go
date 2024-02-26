package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ { // inicia a variavel, vezes, continuidade
		fmt.Printf( "%d ", i)
		time.Sleep(time.Millisecond + 150)
	}
}

func letras() {
	for l := 'a'; l < 'j'; l++ { // inicia a variavel, vezes, continuidade
		fmt.Printf( "%c ", l)
		time.Sleep(time.Millisecond + 170)
	}
}

func main() {
	go numeros() // routines
	go letras()

	time.Sleep( 3 * time.Second) // Espera
}