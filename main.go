package main

import (
	"fmt"
)

// recebe 3 notas e devolve a média
func media(n1, n2, n3 float64) float64 {
	return (n1 + n2 + n3) / 3
}

// recebe a média e devolve o status
func status(m float64) string {
	if m >= 7 {
		return "Aprovado"
	}
	return "Reprovado"
}

// aplica uma função ao resultado de outra
func compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(x A) C {
		return f(g(x))
	}
}

func main() {
	// estrutura para três notas
	type Notas struct {
		n1, n2, n3 float64
	}

	// transformar Notas -> média
	calcMedia := func(n Notas) float64 {
		return media(n.n1, n.n2, n.n3)
	}

	// compor funções: status(media(x))
	resultadoFinal := compose(status, calcMedia)

	// entrada de dados
	var notas Notas
	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&notas.n1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&notas.n2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&notas.n3)

	// execução
	fmt.Println("Resultado:", resultadoFinal(notas))
}
