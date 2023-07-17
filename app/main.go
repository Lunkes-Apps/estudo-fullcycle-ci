package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}
