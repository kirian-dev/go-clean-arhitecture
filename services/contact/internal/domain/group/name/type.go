package name

import "github.com/pkg/errors"

type Name string

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("first name must be greater than or equal to #{MaxLength}")
)

func New(name string) (*Name, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Name(name)
	return &n, nil
}
