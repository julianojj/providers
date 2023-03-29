package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnExceptionIfEmptySubject(t *testing.T) {
	_, err := NewMessage(
		"1",
		&From{
			Name:  "",
			Email: "",
		},
		&To{
			Name:  "",
			Email: "",
		},
		"",
		"Body",
	)
	assert.EqualError(t, err, "Subject cannot be empty")
}
