package storage

import (
	"go-clean-architecture/services/contact/internal/domain/contact"
	"go-clean-architecture/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

type Contact interface {
	CreateContact(c ...*contact.Contact) (*[]contact.Contact, error)
	UpdateContact(c *contact.Contact, ID uuid.UUID) (*contact.Contact, error)
	DeleteContact(ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	GetAllContacts() (*[]contact.Contact, error)
	GetByIDContact(ID uuid.UUID) (*contact.Contact, error)
}

type Group interface {
	CreateGroup(g *group.Group) (*group.Group, error)
	UpdateGroup(g *group.Group, ID uuid.UUID) (*group.Group, error)
	DeleteGroup(ID uuid.UUID) error

	GroupReader
	ContactIntoGroup
}

type GroupReader interface {
	GetAllGroups() (*[]group.Group, error)
	GetByIDGroup(ID uuid.UUID) (*group.Group, error)
	CountGroups() (uint64, error)
}

type ContactIntoGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) (*[]contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactIntoGroup(groupID, contactID uuid.UUID) error
}
