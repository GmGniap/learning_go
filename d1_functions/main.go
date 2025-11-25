package main

import (
	"fmt"
)

func SumTwoInts(a int, b int) int {
	return a + b
}

// Return minimum
func FindMin(a int, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return b
	} else {
		return a
	}
}

// Learning for loop
func simpleForLoop(m int) {
	for p := 1; p <= m; p++ {
		fmt.Printf("P val: %d \n", p)
	}
	fmt.Println("---x---")
}

func main() {
	simpleForLoop(10)

	var x int = 10
	t := SumTwoInts(x, 5)
	min_val := FindMin(x, 10)
	fmt.Printf("Result: %d \n", t)
	if min_val == 0 {
		fmt.Printf("They are equal for num %d \n", x)
	} else {
		fmt.Printf("Min val is %v \n", min_val)
	}
}
