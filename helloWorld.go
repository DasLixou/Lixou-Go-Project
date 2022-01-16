package main

import (
	"fmt"
)

func main() {
	var feeling string

	fmt.Println("Hello World, welcome to Lixou's GoLang Test Project 😄")
	fmt.Print("How do u feel? -> ")
	fmt.Scanf("%s", &feeling)
	fmt.Println("So u are feeling " + feeling + "?")
}
