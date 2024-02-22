package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome  string
	Idade int
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá, meu nome é %s %s e eu tenho %d anos.", p.Nome, p.Sobrenome, p.Idade)
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj PessoaJuridica) String() string {
	return fmt.Sprintf("Olá, nós somos a %s e temos %d anos de existência.", pj.RazaoSocial, pj.Idade)
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("nosso CNPJ é: %s", pj.cnpj)
}

func show(d Document) {
	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).String())
		fmt.Println(d.(PessoaFisica).Doc())
	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).String())
		fmt.Println(d.(PessoaJuridica).Doc())
	default:
		fmt.Println("tipo desconhecido")
	}
}

func main() {
	pf := PessoaFisica{
		Pessoa: Pessoa{
			Nome:  "Tiago",
			Idade: 30,
		},
		Sobrenome: "Santos",
		cpf:       "123.456.789-10",
	}

	pj := PessoaJuridica{
		Pessoa: Pessoa{
			Nome:  "Empresa XYZ",
			Idade: 15,
		},
		RazaoSocial: "Empresa XYZ Ltda.",
		cnpj:        "12.345.678/0001-90",
	}

	show(pf)
	show(pj)
}
