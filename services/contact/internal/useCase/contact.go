package useCase

type ContactUC struct {
	contactRepo Contact
}

func NewContactUC(contactRepo Contact) (*ContactUC, error) {
	return &ContactUC{contactRepo: contactRepo}, nil
}

func (c *ContactUC) Create() {}
