package contact

import (
	"go-clean-architecture/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

func (uc *ContactUC) Create(contacts ...*contact.Contact) (*[]contact.Contact, error) {
	return uc.storage.CreateContact(contacts...)
}
func (uc *ContactUC) Update(contactUpdate *contact.Contact, ID uuid.UUID) (*contact.Contact, error) {
	return uc.storage.UpdateContact(contactUpdate, ID)
}

func (uc *ContactUC) Delete(ID uuid.UUID) error {
	return uc.storage.DeleteContact(ID)
}

func (uc *ContactUC) GetAll() (*[]contact.Contact, error) {
	return uc.storage.GetAllContacts()

}
func (uc *ContactUC) GetByID(ID uuid.UUID) (*contact.Contact, error) {
	return uc.storage.GetByIDContact(ID)
}
