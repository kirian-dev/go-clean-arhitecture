package useCase

import (
	"go-clean-architecture/services/contact/internal/domain/contact"
	"go-clean-architecture/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

type Contact interface {
	Create(c ...*contact.Contact) (*[]contact.Contact, error)
	Update(c *contact.Contact, ID uuid.UUID) (*contact.Contact, error)
	Delete(ID uuid.UUID) error
	ContactReader
}

type ContactReader interface {
	GetAll() (*[]contact.Contact, error)
	GetByID(ID uuid.UUID) (*contact.Contact, error)
}

type Group interface {
	Create(g *group.Group) (*group.Group, error)
	Update(g *group.Group, ID uuid.UUID) (*group.Group, error)
	Delete(ID uuid.UUID) error

	GroupReader
	ContactIntoGroup
}

type GroupReader interface {
	GetAll() (*[]group.Group, error)
	GetByID(ID uuid.UUID) (*group.Group, error)
	Count() (uint64, error)
}

type ContactIntoGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) (*[]contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactIntoGroup(groupID, contactID uuid.UUID) error
}
