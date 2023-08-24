package group

import "go-clean-architecture/services/contact/internal/useCase/adapters/storage"

type GroupUC struct {
	storage storage.Group
	options Options
}

type Options struct{}

func New(storage storage.Group, options Options) *GroupUC {
	var uc = &GroupUC{
		storage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *GroupUC) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
