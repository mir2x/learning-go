package main

import "fmt"

func mapReadWrite() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Linos"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
}
