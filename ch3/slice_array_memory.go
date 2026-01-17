package main

import "fmt"

func sliceArrayMemory() {
	x := [4]int{5,6,7,8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
