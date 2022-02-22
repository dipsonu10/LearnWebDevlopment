package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	n := 0
	return func() int {
		n++
		switch n {
		case 1:
			return 0
		case 2:
			return 1
		default:
			{
				c := a + b
				a = b
				b = c
				return c
			}
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
