package adapters

import (
	"github.com/julianojj/providers/internal/core/domain"
	"github.com/julianojj/providers/internal/core/exceptions"
)

type FakeErrorProvider struct{}

func NewErrorProvider() *FakeErrorProvider {
	return &FakeErrorProvider{}
}

func (f *FakeErrorProvider) Notify(message *domain.Message) (bool, error) {
	return false, exceptions.NewValidationException("Error to send message")
}
