package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	assert.Contains(t, "hello", greeting())
}
