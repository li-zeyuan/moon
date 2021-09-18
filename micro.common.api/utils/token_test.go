package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(320124380004740863, 98, time.Second*3)
	assert.Equal(t, err, nil)
	t.Log(token)
}
