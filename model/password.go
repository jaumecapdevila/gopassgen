package model

// Password struct
type Password struct {
	value []byte
}

// ToString returns the current password value as a string
func (p *Password) ToString() string {
	return string(p.value)
}

// New return a new password from the given random string
func New(value []byte) *Password {
	return &Password{
		value: value,
	}
}
