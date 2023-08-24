package postgres

import (
	"go-clean-architecture/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

func (r *Repository) AddContactToGroup(groupID, contactID uuid.UUID) error {
	panic("not implemented")
}

func (r *Repository) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) (*[]contact.Contact, error) {
	panic("not implemented")
}

func (r *Repository) DeleteContactIntoGroup(groupID, contactID uuid.UUID) error {
	panic("not implemented")
}
