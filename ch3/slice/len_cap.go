package main

import "fmt"

func lenCap() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60, 70, 80)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 90)
	fmt.Println(x, len(x), cap(x))
}
