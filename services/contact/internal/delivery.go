package internal

import "go-clean-architecture/services/contact/internal/domain/contact"

type ContactHandlers interface {
	GetAll() ([]contact.Contact, error)
	GetById(id string) ([]contact.Contact, error)
	Create(c contact.Contact) (contact.Contact, error)
	Update(c contact.Contact, id string) (contact.Contact, error)
	Delete(id string) (contact.Contact, error)
}
