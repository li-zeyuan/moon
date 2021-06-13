package boot

import (
	"context"
	"log"

	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

type Infra struct {
	//Client *redis.Client
	DB      *gorm.DB
	Context context.Context
}

const InfraKey = "infra"

func GetInfra(c context.Context) *Infra {
	if c == nil {
		log.Fatal("content is nil")
		return nil
	}

	infra, ok := c.Value(InfraKey).(*Infra)
	if !ok {
		log.Fatal("can not transfer InfraKey")
		return NewInfra(bson.NewObjectId().Hex())
	}

	return infra
}

func NewInfra(requestID string) *Infra {
	infra := new(Infra)
	infra.DB = config.InitDatabase(&config.Conf)
	return infra
}
