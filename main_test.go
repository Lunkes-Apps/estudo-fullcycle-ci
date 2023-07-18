package main

import "testing"

func TestIfSomaMakeSum(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("O Resultado é inválido: Resultado %d - Experado %d", total, 30)
	}
}
