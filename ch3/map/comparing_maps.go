package main

import (
	"fmt"
	"maps"
)

func comparingMaps() {
	m := map[string]int {
		"hello": 5,
		"world": 0,
	}
	n := map[string]int {
		"world": 0,
		"hello": 5,
	}
	fmt.Println(maps.Equal(m, n))
}
