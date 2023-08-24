package postgres

import (
	"go-clean-architecture/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) (*[]contact.Contact, error) {
	panic("not implemented")
}
func (r *Repository) UpdateContact(contactUpdate *contact.Contact, ID uuid.UUID) (*contact.Contact, error) {
	panic("not implemented")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	panic("not implemented")
}

func (r *Repository) GetAllContacts() (*[]contact.Contact, error) {
	panic("not implemented")
}
func (r *Repository) GetByIDContact(ID uuid.UUID) (*contact.Contact, error) {
	panic("not implemented")
}
