package lastName

import (
	"github.com/pkg/errors"
)

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("last name must be greater than or equal to #{MaxLength}")
)

type LastName string

func (f LastName) String() string {
	return string(f)
}

func New(lastName string) (*LastName, error) {
	if len([]rune(lastName)) > MaxLength {
		return nil, ErrWrongLength
	}
	f := LastName(lastName)
	return &f, nil
}
