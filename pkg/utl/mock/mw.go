package mock

import (
	"github.com/siriusdely/sirius/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*gorsk.User) (string, string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *gorsk.User) (string, string, error) {
	return j.GenerateTokenFn(u)
}
