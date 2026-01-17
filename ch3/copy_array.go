package main

import "fmt"

func copyArray() {
	x := []int{1,2,3,4}
	d := [4]int{5,6,7,8}
	y := make([]int, 2)
	num := copy(y, d[:])
	fmt.Println(y, num)
	num = copy(d[:], x)
	fmt.Println(d, num)
}
