package main

import "fmt"

func stringSlicing() {
	s := "Hello there"
	fmt.Println(s)
	s2 := s[4:7]
	s3 := s[:5]
	s4 := s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
