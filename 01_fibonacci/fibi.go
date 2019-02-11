package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//returns the fabinaci series
func fabinaci(n int) {
	a, b := 0, 1
	if n > 0 {
		fmt.Fprintln(out, a)
	}
	if n >= 1 {
		fmt.Fprintln(out, b)
	}
	for i := 2; i <= n; i++ {
		fmt.Fprintln(out, a+b)
		b = a + b
		a = b - a
	}
}

func main() {
	fabinaci(7)
}
