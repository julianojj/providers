package adapters

import "github.com/julianojj/providers/internal/core/domain"

type Provider interface {
	Notify(message *domain.Message) (bool, error)
}
