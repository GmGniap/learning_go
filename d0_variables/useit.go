package d0_variables

import (
	"fmt"
)

// this is also package block scope
// this function is exported because of identifier "Testing" is capitalized
func Testing() {
	fmt.Println("Planet speed:", planetSpeed)
}
