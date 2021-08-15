package userdbrpc

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/stretchr/testify/assert"
)

func TestGetProfileByPassport(t *testing.T) {
	bInfra := middleware.GetBaseInfra(context.Background(), "")
	pMap, err := GetProfileByPassport(bInfra, []string{"lizeyuan"})
	assert.Equal(t, err, nil)
	t.Log(pMap)
}

func TestInsertProfile(t *testing.T) {
	bInfra := middleware.GetBaseInfra(context.Background(), "")
	pf := new(profile.Profile)
	pf.Name = "lizeyua"
	pf.Passport = "lizeyuan"
	err := CreateProfile(bInfra, []*profile.Profile{pf})
	assert.Equal(t, err, nil)
}

func TestUpdateProfile(t *testing.T) {
	bInfra := middleware.GetBaseInfra(context.Background(), "")
	pf := new(profile.Profile)
	pf.Name = "lizeyua"
	pf.Passport = "lizeyuan"

	err := UpdateProfile(bInfra, []*profile.Profile{pf})
	assert.Equal(t, err, nil)
}
