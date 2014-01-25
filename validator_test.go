package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestDomainValidation(t *testing.T) {
  validator := NewValidator("example.com", "")
  assert.Equal(t, validator("foo@example.com"), true)
  assert.Equal(t, validator("bad@bad.com"), false)
}

func TestEmailListValidation(t *testing.T) {
  validator := NewValidator("", "test_emails.txt")
  assert.Equal(t, validator("one@test.com"), true)
  assert.Equal(t, validator("bad@example.com"), false)
}
