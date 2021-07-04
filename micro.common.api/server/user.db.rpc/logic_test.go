package userdbrpc

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGetProfileByPassport(t *testing.T) {
	bInfra := middleware.GetBaseInfra(context.Background(), "")
	pMap, err := GetProfileByPassport(bInfra, []string{"lizeyuan"})
	assert.Equal(t, err, nil)
	t.Log(pMap)
}
