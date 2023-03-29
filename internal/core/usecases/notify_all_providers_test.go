package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorProvider(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "fake_error_provider"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "Subject",
		Body:    "Body",
	}
	_, err := notifyAllProviders.Execute(input)
	assert.EqualError(t, err, "Error to send message")
}

func TestFailProvider(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "fake_fail_provider"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "Subject",
		Body:    "Body",
	}
	success, _ := notifyAllProviders.Execute(input)
	assert.False(t, success)
}

func TestReturnExceptionIfEmptySubject(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "outlook"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "",
		Body:    "Body",
	}
	_, err := notifyAllProviders.Execute(input)
	assert.EqualError(t, err, "Subject cannot be empty")
}

func TestReturnExceptionIfUnsuportedProvider(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "test"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "",
		Body:    "Body",
	}
	_, err := notifyAllProviders.Execute(input)
	assert.EqualError(t, err, "Subject cannot be empty")
}

func TestNofityWithOutlookProvider(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "outlook"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "Subject",
		Body:    "Body",
	}
	success, _ := notifyAllProviders.Execute(input)
	assert.True(t, success)
}

func TestNofityWithGmailProvider(t *testing.T) {
	var providers []*ProviderInput
	providers = append(providers, &ProviderInput{providerName: "gmail"})
	notifyAllProviders := NewNotifyAllProviders(providers)
	from := &FromInput{
		Name:  "",
		Email: "",
	}
	to := &ToInput{
		Name:  "",
		Email: "",
	}
	input := NotifyAllProvidersInput{
		From:    from,
		To:      to,
		Subject: "Subject",
		Body:    "Body",
	}
	success, _ := notifyAllProviders.Execute(input)
	assert.True(t, success)
}
