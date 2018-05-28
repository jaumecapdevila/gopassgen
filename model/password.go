package model

// Password struct
type Password struct {
	value string
}

// New return a new password from the given random string
func New(value string) *Password {
	return &Password{
		value: value,
	}
}
