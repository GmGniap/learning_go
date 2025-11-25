package main

import (
	"fmt"
)

func main() {
	var items = []int{1, 2, 3, 4, 5}
	// Similar to Python enumerate function
	// i for index and v for value
	for i, v := range items {
		fmt.Println(i, v)
	}
}
