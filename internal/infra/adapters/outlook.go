package adapters

import (
	"github.com/julianojj/providers/internal/core/domain"
)

type Outlook struct{}

func NewOutlook() *Outlook {
	return &Outlook{}
}

func (o *Outlook) Notify(message *domain.Message) (bool, error) {
	return true, nil
}
