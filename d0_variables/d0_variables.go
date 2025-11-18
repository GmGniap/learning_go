package d0_variables

import (
	"fmt"
	"math/rand"
	"time"
)

// this is package block scope
// this var is not exported because of identifier 'planetSpeed' is NOT capitalized
var planetSpeed = 67000

func Calculation() {
	fmt.Println("Hello, Go!")
	fmt.Println("Current time:", time.Now().Format(time.RFC1123))
	fmt.Println("Random number is", rand.Intn(10))
	fmt.Println("Calculated value:", add(10, 32))
}

func add(x int, y int) int {
	return x + y
}
