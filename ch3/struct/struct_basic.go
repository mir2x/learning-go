package main

import "fmt"

type person struct {
	name string
	age int
	pet string
}

func structBasic() {
	julia := person {
		"Julia",
		3,
		"Dog",
	}
	fmt.Println(julia)
	fmt.Println(julia.age)
	julia.pet = "Cat"
	fmt.Println(julia)

	beth := person {
		name: "Beth",
		pet: "Cat",
	}
	fmt.Println(beth)

	var person struct {
		name string
		age int
		pet string
	} 

	person.name = "Tanim"
	person.age = 27
	person.pet = "Cat"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	} {
		name: "Tom",
		kind: "Persian",
	}
	fmt.Println(pet)

}
