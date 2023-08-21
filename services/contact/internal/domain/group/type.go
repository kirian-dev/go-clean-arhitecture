package group

import "github.com/go-playground/validator/v10"

type Group struct {
	ID   string `json:"id" db:"id" validate:"required"`
	Name string `json:"name" db:"name" validate:"required,length=250"`
}

func NewGroup(id, name string) (*Group, error) {
	group := &Group{
		ID:   id,
		Name: name,
	}
	if err := validateGroup(group); err != nil {
		return nil, err
	}

	return group, nil
}

// GetName returns the name of the group.
func (g *Group) GetName() string {
	return g.Name
}

// GetID returns the group's ID.
func (g *Group) GetID() string {
	return g.ID
}
func validateGroup(group *Group) error {
	validate := validator.New()
	if err := validate.Struct(group); err != nil {
		return err
	}

	return nil
}
