package group

import (
	"go-clean-architecture/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

func (uc *GroupUC) Create(group *group.Group) (*group.Group, error) {
	return uc.storage.CreateGroup(group)

}

func (uc *GroupUC) Update(group *group.Group, ID uuid.UUID) (*group.Group, error) {
	return uc.storage.UpdateGroup(group, ID)

}

func (uc *GroupUC) Delete(ID uuid.UUID) error {
	return uc.storage.DeleteGroup(ID)

}

func (uc *GroupUC) GetAll() (*[]group.Group, error) {
	return uc.storage.GetAllGroups()

}

func (uc *GroupUC) GetByID(ID uuid.UUID) (*group.Group, error) {
	return uc.storage.GetByIDGroup(ID)

}

func (uc *GroupUC) Count() (uint64, error) {
	return uc.storage.CountGroups()
}
