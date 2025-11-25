package main

import (
	"fmt"
	"mymodule/d0_variables"
)

// package level declaration
// var actorName string = "Paing"
// var movieName string = "LearnGo"
// var season int = 8

func main() {
	d0_variables.Testing()
	d0_variables.Calculation()

	// Below code will get error for not capitalized
	// fmt.Println("Import var:", d0_variables.planetSpeed)

	var i int

	i = 42
	k := 99
	fmt.Printf("%v, %T \n", k, k)
	fmt.Printf("%v, %T \n", i, i)

	// Can't use := because it isn't new variable. Must use '=' operator.
	k = 1000
	fmt.Printf("Reassigned %d \n", k)

	// Allow to assign k again because of new variable 'name'.
	// But can't change the type of k
	name, k := "Paing", 9000
	fmt.Printf("%s's k is over %d\n", name, k)
}
