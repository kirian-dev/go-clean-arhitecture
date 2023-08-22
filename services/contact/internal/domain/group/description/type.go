package description

import "github.com/pkg/errors"

type Description string

var (
	MaxLength      = 200
	ErrWrongLength = errors.Errorf("first Description must be greater than or equal to #{MaxLength}")
)

func New(description string) (*Description, error) {
	if len([]rune(description)) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Description(description)
	return &n, nil
}
