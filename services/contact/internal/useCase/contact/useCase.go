package contact

import "go-clean-architecture/services/contact/internal/useCase/adapters/storage"

type ContactUC struct {
	storage storage.Contact
	options Options
}

type Options struct{}

func New(storage storage.Contact, options Options) *ContactUC {
	var uc = &ContactUC{
		storage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *ContactUC) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
