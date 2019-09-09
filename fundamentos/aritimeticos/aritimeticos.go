package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise ... bit a bit
	fmt.Println("E =>", a&b)   //0011 & 0010 = 0010
	fmt.Println("OU =>", a|b)  //0011 | 0010 = 0011
	fmt.Println("XOR =>", a^b) //0011 ^ 0010 = 0001

	c := 3.0
	d := 2.0

	// outras operações usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d)) // 3^2
	// math.. vc pode explorar um pouco mais depois

}
