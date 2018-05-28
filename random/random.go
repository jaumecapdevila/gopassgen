package random

import (
	"math/rand"
	"time"

	"github.com/jaumecapdevila/gopassgen/model"
)

// GenerateRandomPasswordService implementation
type GenerateRandomPasswordService struct {
	r *rand.Rand
}

// GenerateFromCharacters returns a new random password generated using the given characters
func (grps *GenerateRandomPasswordService) GenerateFromCharacters(passwordLength int, characters string) *model.Password {
	newPassword := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		newPassword[i] = characters[grps.r.Intn(len(characters))]
	}
	return model.New(newPassword)
}

// New return a new instance of the service
func New() *GenerateRandomPasswordService {
	return &GenerateRandomPasswordService{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
