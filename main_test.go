package main

import "testing"

func TestIfSomaMakeSum(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("O Resultado é inválido: Resultado %d - Experado %d", total, 30)
	}
}

func TestIfSubtracaoIsOk(t *testing.T) {
	total := subtracao(15, 5)
	if total != 10 {
		t.Errorf("O Resultado é inválido: Resultado %d - Experado %d", total, 10)
	}
}
