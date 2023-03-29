package adapters

import (
	"github.com/julianojj/providers/internal/core/domain"
)

type FakeFailProvider struct{}

func NewFakeFailProvider() *FakeFailProvider {
	return &FakeFailProvider{}
}

func (f *FakeFailProvider) Notify(message *domain.Message) (bool, error) {
	return false, nil
}
