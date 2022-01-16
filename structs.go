package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func main() {
	var lixou Person = Person{"Lixou von Baum", 2}
	stackOverflow := Person{"StackOverflow", 13}
	fmt.Println(lixou.name, ",", lixou.age)
	fmt.Println(stackOverflow.name, ",", stackOverflow.age)
}
