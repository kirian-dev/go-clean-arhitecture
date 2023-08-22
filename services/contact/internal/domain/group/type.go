package group

import (
	"go-clean-architecture/services/contact/internal/domain/group/description"
	"go-clean-architecture/services/contact/internal/domain/group/name"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID           uuid.UUID
	Name         name.Name
	Description  description.Description
	CreatedAt    time.Time
	ModifiedAtAt time.Time
	ContactCount uint64
}

func NewWithID(id uuid.UUID, name name.Name, description description.Description, createdAt time.Time, modifiedAt time.Time, contactCount uint64) *Group {
	return &Group{
		ID:           id,
		Name:         name,
		Description:  description,
		CreatedAt:    createdAt,
		ModifiedAtAt: modifiedAt,
		ContactCount: contactCount,
	}
}
func New(name name.Name, description description.Description) *Group {
	var timeNow = time.Now().UTC()
	return &Group{
		ID:           uuid.New(),
		Name:         name,
		Description:  description,
		CreatedAt:    timeNow,
		ModifiedAtAt: timeNow,
	}
}
func (g *Group) GetID() uuid.UUID {
	return g.ID
}

func (g *Group) GetContactCount() uint64 {
	return g.ContactCount
}

func (g *Group) GetDescription() description.Description {
	return g.Description
}

func (g *Group) GetName() name.Name {
	return g.Name
}

func (g *Group) GetCreatedAt() time.Time {
	return g.CreatedAt
}

func (g *Group) GetModifiedAt() time.Time {
	return g.GetModifiedAt()
}
