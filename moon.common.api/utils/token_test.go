package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(318861102280601344, 98, time.Hour*24*30)
	assert.Equal(t, err, nil)
	t.Log(token)
}
