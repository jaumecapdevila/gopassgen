package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Flags
var (
	passwordLength   int
	withSpecialChars bool
	r                *rand.Rand
)

func main() {
	flag.Parse()
	password := generatePassword(passwordLength, withSpecialChars)
	fmt.Printf("Generated password: %s\n", password)
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// Define all the default values for the flags
	flag.IntVar(&passwordLength, "l", 12, " Specify the length of the generated password.")
	flag.BoolVar(&withSpecialChars, "s", true, " Allow special chars in the generated password.")
}

func generatePassword(specifiedLength int, withSpecialChars bool) string {
	const baseChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const specialChars string = "!@#$%^&*()-_=+.?/{}[]"
	availableChars := baseChars
	if withSpecialChars {
		availableChars = baseChars + specialChars
	}
	newPassword := make([]byte, specifiedLength)
	for i := range newPassword {
		newPassword[i] = availableChars[r.Intn(len(availableChars))]
	}
	return string(newPassword)
}

func int2Bool(number int) bool {
	if number != 0 {
		return true
	}
	return false
}
