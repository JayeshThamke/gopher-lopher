package main

import (
	"fmt"

	"myhome.io/pkg/greet"
)

func main() {
	fmt.Println(greet.Greet("Jayesh"))
	fmt.Println(greet.GoodBye("Jayesh"))
}
