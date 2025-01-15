package main

import (
	"testing"
)

// func TestFactorial(t *testing.T) {
// 	expected := big.NewInt(24)

// 	result := Factorial(4)

// 	if result.Int64() != expected.Int64() {
// 		t.Error("not equal")
// 	}
// }

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(400)
	}
}
