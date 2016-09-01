package moretypes

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a := 0
	b := 1
	return func() int {
     	a0 := a
		a = b
		b = b + a0
		return a0
	}
}

func s26(n int) {
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}
