package dao

import (
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

func (d *UserDao) Update(infra *middleware.Infra, models []*inner.UserProfileModel) error {
	if len(models) == 0 {
		return nil
	}

	for _, m := range models {
		if m.Uid <= 0 {
			continue
		}

		values := make(map[string]interface{})
		if len(m.Name) > 0 {
			values[inner.ColumnName] = m.Name
		}
		if len(m.Name) > 0 {
			values[inner.ColumnPassport] = m.Passport
		}
		if len(m.Name) > 0 {
			values[inner.ColumnPassword] = m.Password
		}
		if len(m.Name) > 0 {
			values[inner.ColumnGender] = m.Gender
		}
		if len(m.Name) > 0 {
			values[inner.ColumnBirth] = m.Birth
		}
		if len(m.Name) > 0 {
			values[inner.ColumnPortrait] = m.Portrait
		}
		if len(m.Name) > 0 {
			values[inner.ColumnHometown] = m.Hometown
		}
		if len(m.Name) > 0 {
			values[inner.ColumnDescription] = m.Description
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
		First(m).Error
	if err != nil {
		infra.Log.Error("get user by openid error: ", err)
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
