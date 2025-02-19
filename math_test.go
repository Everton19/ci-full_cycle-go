package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado %d, resultado obtido %d", 30, total)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 15)

	if total != 0 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado %d, resultado obtido %d", 0, total)
	}
}

func TestDiv(t *testing.T) {
	total := Div(15, 15)

	if total != 1 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado %d, resultado obtido %d", 1, total)
	}
}

func TestMult(t *testing.T) {
	total := Mult(15, 15)

	if total != 225 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado %d, resultado obtido %d", 225, total)
	}
}
