package main

import (
	"fmt"
	"lunkes/calc/internal/entity"
)

func main() {
	fmt.Println(soma(10, 10))
	order, err := entity.NewOrder("test", 3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(order.FinalPrice)
}

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}
