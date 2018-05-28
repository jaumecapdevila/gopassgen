package model

// GeneratePasswordService interface
type GeneratePasswordService interface {
	GenerateFromCharacters(int, string) *Password
}
