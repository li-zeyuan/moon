package dao

import (
	"github.com/li-zeyuan/micro/moon.common.api/utils"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/app/model/inner"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDao {
	return &UserDao{
		db,
	}
}

func (d *UserDao) Save(infra *middleware.Infra, models []*inner.UserProfileModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.TableNameUserProfile).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create users error: ", err)
		return err
	}

	return nil
}

func (d *UserDao) Update(infra *middleware.Infra, models []*model.UserProfileUpdateField) error {
	if len(models) == 0 {
		return nil
	}

	for _, m := range models {
		if m.Uid <= 0 {
			continue
		}

		values := make(map[string]interface{})
		if m.Name != nil {
			values[inner.ColumnName] = *m.Name
		}
		if m.Password != nil {
			values[inner.ColumnPassword] = *m.Password
		}
		if m.Gender != nil {
			values[inner.ColumnGender] = *m.Gender
		}
		if m.Birth != nil {
			values[inner.ColumnBirth] = utils.TimeStamp2Time(*m.Birth)
		}
		if m.Portrait != nil {
			values[inner.ColumnPortrait] = *m.Portrait
		}
		if m.Hometown != nil {
			values[inner.ColumnHometown] = *m.Hometown
		}
		if m.Phone != nil {
			values[inner.ColumnPhone] = *m.Phone
		}

		err := d.db.Table(inner.TableNameUserProfile).
			Where(inner.ColumnUid+" = ?", m.Uid).
			Updates(&values).Error
		if err != nil {
			infra.Log.Errorf("update profile by uid:  error: %v", m.Uid, err)
			return err
		}
	}

	return nil
}

func (d *UserDao) GetByOpenid(infra *middleware.Infra, openid string) (*inner.UserProfileModel, error) {
	m := new(inner.UserProfileModel)
	err := d.db.Table(inner.TableNameUserProfile).
		Where("openid = ?", openid).
		Where("deleted_at is null").
		First(m).Error
	if err != nil {
		infra.Log.Error("get user by openid error: ", err)
		return nil, err
	}

	return m, nil
}

func (d *UserDao) GetOne(infra *middleware.Infra, uid int64) (*inner.UserProfileModel, error) {
	m := new(inner.UserProfileModel)
	err := d.db.Table(inner.TableNameUserProfile).
		Where("uid = ?", uid).
		Where("deleted_at is null").
		First(m).Error
	if err != nil {
		infra.Log.Error("get user by uid error: ", err)
		return nil, err
	}

	return m, nil
}

func (d *UserDao) Get(infra *middleware.Infra, uids []int64) ([]*inner.UserProfileModel, error) {
	if len(uids) == 0 {
		return nil, nil
	}

	models := make([]*inner.UserProfileModel, 0)
	err := d.db.Table(inner.TableNameUserProfile).
		Where("").
		Find(&models).Error
	if err != nil {
		infra.Log.Error("get user by uids error: ", err)
		return nil, err
	}

	return models, nil
}

func (d *UserDao) GetByPassport(infra *middleware.Infra, passports []string) ([]*inner.UserProfileModel, error) {
	if len(passports) == 0 {
		return nil, nil
	}

	models := make([]*inner.UserProfileModel, 0)
	err := d.db.Table(inner.TableNameUserProfile).
		Where("passport in (?)", passports).
		Find(&models).Error
	if err != nil {
		infra.Log.Error("get user by passport error: ", err)
		return nil, err
	}

	return models, nil
}
