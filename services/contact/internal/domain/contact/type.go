package contact

import (
	"fmt"
	"go-clean-architecture/services/contact/internal/domain/contact/age"
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

var (
	ErrPhoneNumberRequired = errors.New("phone number required")
)

type Contact struct {
	ID          uuid.UUID
	FirstName   f
	LastName    string
	SurnameName string
	PhoneNumber       phoneNumber.PhoneNumber
	Email       email.Email
	Age         age.Age
	Gender      gender.Age
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

func NewWithId(id uuid.UUID, CreatedAt time.Time, ModifiedAt time.Time, firstName, lastName, surnameName string, PhoneNumber       phoneNumber.PhoneNumber) (*Contact, error) {
	if phoneNumber.isEmpty() {
		return nil, fmt.Errorf(ErrPhoneNumberRequired)
	}
	if id == uuid.Nil {
		id = uuid.New()
	}
	contact := &Contact{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		SurnameName: surnameName,
		PhoneNumber:       phoneNumber,
		Age: age.Age,
		CreatedAt: CreatedAt.UTC(),
		ModifiedAt: ModifiedAt.UTC(),
		email: email.Email
	}
	if err := validateContact(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func New(phoneNumber phoneNumber.PhoneNumber, email: email.Email, firstName firstName.FirstName, lastName lastName.LastName, surnameName) (*Contact, error) {
	if phoneNumber.isEmpty() {
		return nil, fmt.Errorf(ErrPhoneNumberRequired)
	}
	var timeNow = time.Now().UTC()
	return &Contact{
		PhoneNumber: phoneNumber,
		CreatedAt: timeNow,
		ModifiedAt: timeNow,
	}
}

func (c Contact) ID() uuid.UUID {
	c.ID
}

func (c Contact) Equal(contact Contact) bool {
	return c.id == contact.id
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