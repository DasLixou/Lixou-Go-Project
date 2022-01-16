package main

import (
	"fmt"
)

func greet(name string) string {
	return "Howdy, " + name
}

func askForFeeling() string {
	var feeling string

	fmt.Print("How do u feel? -> ")
	fmt.Scanf("%s", &feeling)

	return feeling
}

func main() {
	fmt.Println("Hello World, welcome to Lixou's GoLang Test Project 😄")
	fmt.Println(greet("Lixou"))
	feeling := askForFeeling()
	fmt.Println("So u are feeling " + feeling + "?")
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)
}
