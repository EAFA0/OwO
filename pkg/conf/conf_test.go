package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	ProjectPathSkipStep = 3
	Init()

	assert.Equal(t, "dev", GetString("env_name"))
}
