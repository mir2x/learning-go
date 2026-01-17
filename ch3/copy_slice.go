package main

import "fmt"

func copySlice() {
	x := []int{1,2,3,4,5}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)
	fmt.Println(cap(x))
	fmt.Println(cap(y))
	y = append(y, 6)
	fmt.Println(x)
	fmt.Println(y)
	a := []int{1,2,3,4}
	b := make([]int, 2)
	num = copy(b, a[:2])
	fmt.Println(b, num)	
	num = copy(a[:3], a[1:])
	fmt.Println(a, num)
} 

