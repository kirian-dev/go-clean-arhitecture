package age

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	MaxLength      uint64 = 200
	ErrWrongLength        = errors.Errorf("age must be greater than or equal to #{MaxLength}")
)

type Age uint8

func (a Age) String() string {
	return strconv.FormatUint(uint64(a), 10)
}

func New(age uint64) (*Age, error) {
	if age > MaxLength {
		return nil, ErrWrongLength
	}
	a := Age(age)
	return &a, nil
}

func (a Age) Equal(age Age) bool {
	return a == age
}

func (a Age) Less(age Age) bool {
	return a < age
}

func (a Age) Greater(age Age) bool {
	return a > age
}
