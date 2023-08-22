package phoneNumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	regexpPhoneNumber    = regexp.MustCompile(`^\+\d{1,3}-\d{4,14}$`)
	ErrFormatPhoneNumber = errors.New("Invalid phone number format")
)

type PhoneNumber struct {
	value string
}

func New(phoneNumber string) (PhoneNumber, error) {
	if len(phoneNumber) == 0 {
		return PhoneNumber{}, nil
	}

	if !regexpPhoneNumber.MatchString(phoneNumber) {
		return PhoneNumber{}, ErrFormatPhoneNumber
	}

	return PhoneNumber{value: phoneNumber}, nil
}

func (p PhoneNumber) IsEmpty() bool {
	return len(strings.TrimSpace(p.String())) == 0
}

func (p PhoneNumber) String() string {
	return p.value
}

func (p PhoneNumber) Equal(phoneNumber PhoneNumber) bool {
	return p.value == phoneNumber.value
}

func (p PhoneNumber) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, p.value)), nil
}

func (p *PhoneNumber) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	if !regexpPhoneNumber.MatchString(str) {
		return ErrFormatPhoneNumber
	}

	p.value = str
	return nil
}
