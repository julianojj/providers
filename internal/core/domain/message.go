package domain

import "github.com/julianojj/providers/internal/core/exceptions"

type From struct {
	Name  string
	Email string
}

type To struct {
	Name  string
	Email string
}

type Message struct {
	ID      string
	From    *From
	To      *To
	Subject string
	Body    string
}

func NewMessage(
	id string,
	from *From,
	to *To,
	subject string,
	body string,
) (*Message, error) {
	message := &Message{
		ID:      id,
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
	err := message.Validate()
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (m *Message) Validate() error {
	if m.Subject == "" {
		return exceptions.NewValidationException("Subject cannot be empty")
	}
	return nil
}
