package main

import (
	"math/big"
	"net/http"
	_ "net/http/pprof" // register pprof handlers
)

func main() {
	http.HandleFunc("/fact", func(w http.ResponseWriter, r *http.Request) {
		res := Factorial(100)
		w.Write([]byte(res.String()))
	})
	http.ListenAndServe(":3000", nil)
}

func Factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	f := big.NewInt(n)
	return f.Mul(f, Factorial(n-1))
}
