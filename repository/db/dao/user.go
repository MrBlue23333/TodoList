package dao

import (
	"context"
	"demo/repository/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	if c == nil {
		c = context.Background()
	}
	return &UserDao{NewDbClient(c)}
}

func (u *UserDao) FindUserByName(userName string) (user *model.UserModel, err error) {
	err = u.db.Model(&model.UserModel{}).
		Where("user_name = ?", userName).
		First(&user).Error
	return
}

func (u *UserDao) FindUserByUid(uid int64) (user *model.UserModel, err error) {
	err = u.db.Model(&model.UserModel{}).
		Where("id = ?", uid).
		First(&user).Error
	return
}

func (u *UserDao) Create(user *model.UserModel) error {
	return u.db.Create(user).Error
}
