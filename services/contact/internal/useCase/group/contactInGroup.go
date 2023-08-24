package group

import (
	"go-clean-architecture/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

func (uc *GroupUC) AddContactToGroup(groupID, contactID uuid.UUID) error {
	return uc.storage.AddContactToGroup(groupID, contactID)

}

func (uc *GroupUC) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) (*[]contact.Contact, error) {
	return uc.storage.CreateContactIntoGroup(groupID, contacts...)
}

func (uc *GroupUC) DeleteContactIntoGroup(groupID, contactID uuid.UUID) error {
	return uc.storage.DeleteContactIntoGroup(groupID, contactID)

}
