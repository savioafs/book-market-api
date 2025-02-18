package utils

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	as := assert.New(t)

	id, err := NewID()
	as.Nil(err)

	err = uuid.Validate(id)
	as.Nil(err)

	as.NotNil(id)
}

func TestNewCode(t *testing.T) {
	as := assert.New(t)

	code, err := NewCode("marcos")
	as.Nil(err)

	as.NotNil(code)
}
