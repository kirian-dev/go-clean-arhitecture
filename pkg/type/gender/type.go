package gender

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	regexpGender     = regexp.MustCompile(`^(male|female|other)$`)
	ErrInvalidGender = errors.New("gender must be one of male, female or other")
)

type Gender struct {
	value string
}

func New(gender string) (Gender, error) {
	if len(gender) == 0 {
		return Gender{}, nil
	}

	if !regexpGender.MatchString(strings.ToLower(gender)) {
		return Gender{}, ErrInvalidGender
	}

	return Gender{value: gender}, nil
}

func (g Gender) String() string {
	return g.value
}

func (g Gender) Equal(gender Gender) bool {
	return strings.ToLower(g.value) == strings.ToLower(gender.value)
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, g.value)), nil
}

func (g *Gender) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	if !regexpGender.MatchString(strings.ToLower(str)) {
		return ErrInvalidGender
	}

	g.value = str
	return nil
}
