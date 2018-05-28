package main

import (
	"flag"
	"fmt"

	"github.com/jaumecapdevila/gopassgen/random"
)

// Flags
var (
	passwordLength   int
	withSpecialChars bool
)

func init() {
	flag.IntVar(&passwordLength, "l", 12, " Specify the length of the generated password.")
	flag.BoolVar(&withSpecialChars, "s", true, " Allow special chars in the generated password.")
}

func main() {
	flag.Parse()
	service := random.New()
	password := service.GenerateFromCharacters(passwordLength, getCharactersToUse())
	fmt.Printf("Generated password: %s\n", password.ToString())
}

func getCharactersToUse() string {
	const baseChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	const specialChars string = "!@#$%^&*()-_=+.?/{}[]"

	if !withSpecialChars {
		return baseChars
	}
	return baseChars + specialChars
}
