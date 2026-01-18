package main

import "fmt"

func commaOkIdiom() {
	m := map[string]int {
		"Hello": 5,
		"world": 0,
	}
	fmt.Println(m)
	v, ok := m["Hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["Goodbye"]
	fmt.Println(v, ok)
	m["Goodbye"]++
	v, ok = m["Goodbye"]
	fmt.Println(v, ok)
	fmt.Println(m)
	delete(m, "Hello")
	fmt.Println(m, len(m))
	clear(m)
	fmt.Println(m, len(m))
}
