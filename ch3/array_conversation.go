package main

import "fmt"

func arrayConversation() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
	xArrayPointer := (*[4]int)(xSlice)
	fmt.Println(*xArrayPointer)
	xSlice[1] = 20
	xArrayPointer[2] = 30
	fmt.Println(xSlice)
	fmt.Println(*xArrayPointer)
}
