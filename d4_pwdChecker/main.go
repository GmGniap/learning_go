package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	// Convert to unicode format
	unicode_pw := []rune(pw)

	if len(unicode_pw) < 8 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasNum := false
	hasSymb := false

	for _, val := range unicode_pw {
		if unicode.IsLower(val) {
			hasLower = true
		}
		if unicode.IsUpper(val) {
			hasUpper = true
		}
		if unicode.IsNumber(val) {
			hasNum = true
		}
		if unicode.IsSymbol(val) || unicode.IsPunct(val) {
			hasSymb = true
		}
	}
	return hasLower && hasUpper && hasNum && hasSymb
}

func main() {
	if passwordChecker("This!I5A") {
		fmt.Println("Good password!")
	} else {
		fmt.Println("Bad password.")
	}
}
