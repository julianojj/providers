package factory

import (
	"github.com/julianojj/providers/internal/core/exceptions"
	"github.com/julianojj/providers/internal/infra/adapters"
)

type ProviderFactory struct{}

func NewProviderFactory() *ProviderFactory {
	return &ProviderFactory{}
}

func (f *ProviderFactory) CreateProvider(providerName string) (adapters.Provider, error) {
	if providerName == "gmail" {
		return adapters.NewGmail(), nil
	}
	if providerName == "outlook" {
		return adapters.NewOutlook(), nil
	}
	if providerName == "fake_error_provider" {
		return adapters.NewErrorProvider(), nil
	}
	if providerName == "fake_fail_provider" {
		return adapters.NewFakeFailProvider(), nil
	}
	return nil, exceptions.NewValidationException("Unsupported provider")
}
