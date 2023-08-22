package contact

import (
	"errors"
	"fmt"
	"go-clean-architecture/pkg/type/email"
	"go-clean-architecture/pkg/type/gender"
	phoneNumber "go-clean-architecture/pkg/type/phone"
	"go-clean-architecture/services/contact/internal/domain/contact/age"
	"go-clean-architecture/services/contact/internal/domain/contact/firstName"
	"go-clean-architecture/services/contact/internal/domain/contact/lastName"
	"go-clean-architecture/services/contact/internal/domain/contact/surname"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPhoneNumberRequired = errors.New("phone number required")
)

type Contact struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	ModifiedAt time.Time

	FirstName   firstName.FirstName
	LastName    lastName.LastName
	SurnameName surname.Surname

	PhoneNumber phoneNumber.PhoneNumber
	Email       email.Email

	Age    age.Age
	Gender gender.Gender
}

func NewWithId(id uuid.UUID, createdAt time.Time, modifiedAt time.Time, firstName firstName.FirstName, lastName lastName.LastName, surnameName surname.Surname, phoneNumber phoneNumber.PhoneNumber, age age.Age) (*Contact, error) {
	if phoneNumber.IsEmpty() {
		return nil, fmt.Errorf("%w", ErrPhoneNumberRequired)
	}
	if id == uuid.Nil {
		id = uuid.New()
	}
	contact := &Contact{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		SurnameName: surnameName,
		PhoneNumber: phoneNumber,
		Age:         age,
		CreatedAt:   createdAt.UTC(),
		ModifiedAt:  modifiedAt.UTC(),
		Email:       email.Email{},
	}

	return contact, nil
}

func New(phoneNumber phoneNumber.PhoneNumber, firstName firstName.FirstName, lastName lastName.LastName, surnameName surname.Surname, age age.Age, email email.Email) (*Contact, error) {
	if phoneNumber.IsEmpty() {
		return nil, fmt.Errorf("%w", ErrPhoneNumberRequired)
	}
	var timeNow = time.Now().UTC()
	return &Contact{
		PhoneNumber: phoneNumber,
		CreatedAt:   timeNow,
		ModifiedAt:  timeNow,
		Age:         age,
		Email:       email,
	}, nil
}

func (c Contact) GetID() uuid.UUID {
	return c.ID
}

func (c Contact) Equal(contact Contact) bool {
	return c.ID == contact.ID
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.FirstName, c.LastName, c.SurnameName)
}

func (c Contact) GetAge() age.Age {
	return c.Age
}

func (c Contact) GetPhoneNumber() phoneNumber.PhoneNumber {
	return c.PhoneNumber
}

func (c Contact) GetEmail() email.Email {
	return c.Email
}

func (c Contact) GetGender() gender.Gender {
	return c.Gender
}

func (c Contact) GetCreatedAt() time.Time {
	return c.CreatedAt
}

// Getter for ModifiedAt
func (c Contact) GetModifiedAt() time.Time {
	return c.ModifiedAt
}
