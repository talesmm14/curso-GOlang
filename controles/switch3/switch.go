package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.4))
	fmt.Println(tipo(2))
	fmt.Println(tipo("tipo"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
