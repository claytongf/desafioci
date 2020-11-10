package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	result := math.Sqrt(64)
	var response float64 = 8

	if result != response {
		t.Errorf("Resultado Inválido! retornou %f, requerido %f", result, response)
	}
}
