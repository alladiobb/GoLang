package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)
	if total != 30 {
		t.Error("Resultado é errado, valor esperado era ")
	}
}
