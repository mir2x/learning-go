package main

import "fmt"

func fullSliceExpression() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(y)
	fmt.Println(z)
	y = append(y, "x", "y")
	z = append(z, "p", "q")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
