package internal

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/boot"
)

func TestMockDao_Save(t *testing.T) {
	infra := boot.NewInfra("")
	mockDao := NewMock(infra.DB)

	batch := make([]*inner.MockModel, 0)
	for i := 0; i < 10000000; i++ {
		mock := new(inner.MockModel)
		mock.Field1 = rand.Int63n(98999) + 1000
		mock.Field2 = rand.Int63n(98999) + 1000
		mock.Field3 = rand.Int63n(98999) + 1000
		mock.Field4 = rand.Int63n(98999) + 1000
		mock.Field5 = rand.Int63n(98999) + 1000
		batch = append(batch, mock)

		if len(batch) == 500 {
			if err := mockDao.Save(batch); err != nil {
				return
			}
			batch = make([]*inner.MockModel, 0)
			log.Println(i)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func TestRand(t *testing.T) {
	fmt.Println(rand.Int63())
}
