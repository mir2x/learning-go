package main

import "fmt"

func stringToSlice() {
	s := "Hello 🌞"
	bs := []byte(s)
	rs := []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
