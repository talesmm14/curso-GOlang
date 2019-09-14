package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	// Nâo tem :(
	//return nota >= 6 ? "Aprovado" : "Reprovado"
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.3))
}
