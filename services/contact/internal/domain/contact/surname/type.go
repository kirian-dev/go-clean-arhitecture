package surname

import (
	"github.com/pkg/errors"
)

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("last name must be greater than or equal to #{MaxLength}")
)

type Surname string

func (f Surname) String() string {
	return string(f)
}

func New(surname string) (*Surname, error) {
	if len([]rune(surname)) > MaxLength {
		return nil, ErrWrongLength
	}
	f := Surname(surname)
	return &f, nil
}
