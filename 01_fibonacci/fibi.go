package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//prints the fibonacci series
func fibonacci(n int) {
	a, b := 1, 1
	if n >=1 {
		for i := 1; i <= n; i++ {
			fmt.Fprintln(out, a)
			a, b = b, (a + b)
		}
  	} else {
		 Log.e 
	  }
}

func main() {
	fibonacci(7)
}
