package adapters

import (
	"github.com/julianojj/providers/internal/core/domain"
)

type Gmail struct{}

func NewGmail() *Gmail {
	return &Gmail{}
}

func (g *Gmail) Notify(message *domain.Message) (bool, error) {
	return true, nil
}
