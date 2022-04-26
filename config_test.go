package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigGetClient(t *testing.T) {
	client, err := GetClient()

	assert.NoError(t, err)
	assert.NotEqual(t, client, nil)
}

func TestConfigUpdateStateFromKeyBindNextConvo(t *testing.T) {
	state := MakeMockState()
	assert.Equal(t, state.pos, 0)
	newState := UpdateStateFromKeyBind(state, 'l')
	assert.Equal(t, newState.pos, 1)
}

func TestConfigUpdateStateFromKeyBindNextMessage(t *testing.T) {
	state := MakeMockState()
	state = UpdateStateFromKeyBind(state, 'j')
	assert.Equal(t, state.conversations[state.pos].messagePos, 3)
	state = UpdateStateFromKeyBind(state, 'k')
	assert.Equal(t, state.conversations[state.pos].messagePos, 2)
	state = UpdateStateFromKeyBind(state, 'k')
	assert.Equal(t, state.conversations[state.pos].messagePos, 1)
}