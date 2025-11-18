package main

import (
	"fmt"
)

// package level declaration
var actorName string = "Paing"
var movieName string = "LearnGo"
var season int = 8

func main() {
	var i int
	i = 42
	k := 99
	fmt.Printf("%v, %T \n", k, k)
	fmt.Printf("%v, %T \n", i, i)
}
