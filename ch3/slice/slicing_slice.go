package main

import "fmt"
func slicingSlice() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	x[1] = "y"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	y[0] = "x"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
