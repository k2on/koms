package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientProviderNone(t *testing.T) {
	client, err := NewClient([]Provider{}, nil)
	assert.NoError(t, err, "New client should not return an error")

	providers := client.GetProviders()
	assert.Empty(t, providers, "New client with no providers should return no providers")
}

func TestClientProviderMockA(t *testing.T) {
	provider, _ := NewProviderMockA()
	client, err := NewClient([]Provider{provider}, nil)
	assert.NoError(t, err, "New client w/ mock provider should not return an error")

	providers := client.GetProviders()
	assert.Equal(t, len(providers), 1)
}

func TestClientContactsMockA(t *testing.T) {
	contacts, _ := NewContactsMock()
	_, err := NewClient([]Provider{}, contacts)

	assert.NoError(t, err, "New client with mock contacts should not return an error")
}

func TestClientProviderMockAGetConversations(t *testing.T) {
	provider, _ := NewProviderMockA()
	client, _ := NewClient([]Provider{provider}, nil)

	conversations, err := client.GetConversations()
	assert.NoError(t, err)

	assert.Equal(t, len(conversations), 2)
}

func TestClientContactsMockAGetContact(t *testing.T) {
	contacts, _ := NewContactsMock()
	client, _ := NewClient([]Provider{}, contacts)

	contact, err := client.GetContact("0")
	assert.NoError(t, err)

	assert.Equal(t, contact.id, "0")
}

func TestClientMockABContact(t *testing.T) {
	contacts, _ := NewContactsMock()
	providerA, _ := NewProviderMockA()
	providerB, _ := NewProviderMockB()
	client, _ := NewClient([]Provider{providerA, providerB}, contacts)

	conversations, err := client.GetConversations()
	assert.NoError(t, err)
	assert.Equal(t, len(conversations), 2)

}