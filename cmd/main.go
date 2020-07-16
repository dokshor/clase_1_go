package main

import (
	calculator "calculator/src"
	"fmt"
)

var VERSION = 1.0

func main() {
	var a int = 10
	var b int = 5

	fmt.Printf("Calculadora Version %v\n\n", VERSION)
	fmt.Printf("a=%v , b=%v, a+b=%v\n", a, b, calculator.Sumar(a, b))
	fmt.Printf("a=%v , b=%v, a-b=%v\n", a, b, calculator.Restar(a, b))
	fmt.Printf("a=%v , b=%v, a/b=%v\n", a, b, calculator.Division(a, b))

}
