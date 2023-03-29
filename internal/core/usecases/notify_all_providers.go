package usecases

import (
	"github.com/google/uuid"
	"github.com/julianojj/providers/internal/core/domain"
	"github.com/julianojj/providers/internal/core/factory"
)

type NotifyAllProviders struct {
	Providers []*ProviderInput
}

type ProviderInput struct {
	providerName string
}

type FromInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ToInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NotifyAllProvidersInput struct {
	From    *FromInput `json:"from"`
	To      *ToInput   `json:"to"`
	Subject string     `json:"subject"`
	Body    string     `json:"body"`
}

func NewNotifyAllProviders(providers []*ProviderInput) *NotifyAllProviders {
	return &NotifyAllProviders{
		Providers: providers,
	}
}

func (n *NotifyAllProviders) Execute(input NotifyAllProvidersInput) (bool, error) {
	from := &domain.From{
		Name:  input.From.Name,
		Email: input.From.Email,
	}
	to := &domain.To{
		Name:  input.To.Name,
		Email: input.To.Email,
	}
	message, err := domain.NewMessage(
		uuid.NewString(),
		from,
		to,
		input.Subject,
		input.Body,
	)
	if err != nil {
		return false, err
	}
	for _, providerInput := range n.Providers {
		providerFactory := factory.NewProviderFactory()
		provider, err := providerFactory.CreateProvider(providerInput.providerName)
		if err != nil {
			return false, err
		}
		success, err := provider.Notify(message)
		if err != nil {
			return false, err
		}
		if success {
			return true, nil
		}
	}
	return false, nil
}
