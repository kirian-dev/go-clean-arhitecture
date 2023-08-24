package postgres

import (
	"go-clean-architecture/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("not implemented")
}

func (r *Repository) UpdateGroup(groupUpdate *group.Group, ID uuid.UUID) (*group.Group, error) {
	panic("not implemented")
}

func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	panic("not implemented")
}

func (r *Repository) GetAllGroups() (*[]group.Group, error) {
	panic("not implemented")
}

func (r *Repository) GetByIDGroup(ID uuid.UUID) (*group.Group, error) {
	panic("not implemented")
}

func (r *Repository) CountGroups() (uint64, error) {
	panic("not implemented")
}
