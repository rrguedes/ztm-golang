//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type IdLivro int
type IdMembro int

type Livro struct {
	titulo      string
	quantidade  int
	emprestados int
}

type Membro struct {
	nome        string
	emprestimos map[IdLivro]Emprestimo
}

type Emprestimo struct {
	retirada  time.Time
	devolucao time.Time
}

type Biblioteca struct {
	livrosDisponiveis map[IdLivro]Livro
	membros           map[IdMembro]Membro
}

func emprestarLivro(biblioteca *Biblioteca, idLivro IdLivro, membro *Membro) bool {
	livro, encontrado := biblioteca.livrosDisponiveis[idLivro]
	if !encontrado {
		fmt.Println("O livro solicitado não existe.")
		return false
	}
	if livro.quantidade <= livro.emprestados {
		fmt.Println("No momento não existem exemplares disponíveis deste título.")
		return false
	}
	livro.emprestados += 1
	membro.emprestimos[idLivro] = Emprestimo{retirada: time.Now()}
	return true
}

func devolverLivro(biblioteca *Biblioteca, idLivro IdLivro, membro *Membro) bool {
	livro, encontrado := biblioteca.livrosDisponiveis[idLivro]
	if !encontrado {
		fmt.Println("O livro solicitado não existe.")
		return false
	}
	temRetirada := !membro.emprestimos[idLivro].retirada.IsZero()
	if !temRetirada {
		fmt.Println("No momento não existem exemplares disponíveis deste título.")
		return false
	}
	livro.emprestados -= 1
	membro.emprestimos[idLivro] = Emprestimo{devolucao: time.Now()}
	return true
}

func main() {
	biblioteca := Biblioteca{
		livrosDisponiveis: make(map[IdLivro]Livro),
		membros:           make(map[IdMembro]Membro),
	}
	biblioteca.livrosDisponiveis[1] = Livro{titulo: "O Segredo", quantidade: 5}
	biblioteca.livrosDisponiveis[2] = Livro{titulo: "Código Da Vinci", quantidade: 10}
	biblioteca.livrosDisponiveis[3] = Livro{titulo: "Aprendendo GO", quantidade: 10}
	biblioteca.livrosDisponiveis[4] = Livro{titulo: "Delphi Segredos da Linguagem", quantidade: 10}

	biblioteca.membros[1] = Membro{nome: "Rafael Reis Guedes"}
	biblioteca.membros[2] = Membro{nome: "Cibelle Guedes"}
	biblioteca.membros[3] = Membro{nome: "Gabriel Guedes"}
	biblioteca.membros[4] = Membro{nome: "Gustavo Guedes"}

	printLivrosDisponiveis(&biblioteca)

	member := biblioteca.membros[1]
	emprestarLivro(&biblioteca, 1, &member)
}

func printLivrosDisponiveis(biblioteca *Biblioteca) {
	for _, livro := range biblioteca.livrosDisponiveis {
		printStatusLivro(&livro)
	}
}

func printStatusLivro(livro *Livro) {
	fmt.Println("Livro:", livro.titulo, "Qtd:", livro.quantidade, "Disponíveis:", livro.quantidade-livro.emprestados)
}
