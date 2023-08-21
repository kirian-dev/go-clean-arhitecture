package firstName

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	MaxLength      = 50
	ErrWrongLength        = errors.Errorf("first name must be greater than or equal to #{MaxLength}")
)

type FirstName string

func (f FirstName) String() string {
	return string(f)
}

func New(firstName string) (*FirstName, error) {
	if len([]rune(firstName)) > MaxLength {
		nil, ErrWrongLength
	}
	f := FirstName(firstName)
	return &f, nil
}
