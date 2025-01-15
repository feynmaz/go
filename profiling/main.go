package main

import "math/big"

func main() {

}

func Factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	f := big.NewInt(n)
	return f.Mul(f, Factorial(n-1))
}
