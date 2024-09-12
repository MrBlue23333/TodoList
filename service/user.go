package service

import (
	"context"
	"demo/pkg/ctl"
	"demo/pkg/utils"
	"demo/repository/db/dao"
	"demo/repository/db/model"
	"demo/types"
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

// GetUserSrv 通过Once完成小型结构体的实例化，而不需要显式地写在其他地方
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (*UserSrv) UserRegister(c context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(c)
	u, err := userDao.FindUserByName(req.UserName)
	switch err {
	case gorm.ErrRecordNotFound:
		{
			u = &model.UserModel{
				UserName: req.UserName,
			}
			if err = u.SetPassword(req.Password); err != nil {
				utils.LogrusObj.Error(err)
				return
			}
			if err = userDao.Create(u); err != nil {
				utils.LogrusObj.Error(err)
				return
			}

			return ctl.RespSuccess(), nil
		}
	case nil:
		{
			err = errors.New("user already exists")
			utils.LogrusObj.Error(err)
			return
		}
	default:
		return
	}
}

func (*UserSrv) UserLogin(c context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByName(req.UserName)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	//检查密码
	if !user.CheckPassword(req.Password) {
		err = errors.New("username or password error")
		utils.LogrusObj.Error(err)
		return
	}

	//jwt签发token
	token, err := utils.GenerateToken(user.Id, user.UserName)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}

	//返回携带token的loginResp
	userResp := types.TokenData{
		User: types.UserInfoResp{
			Id:       user.Id,
			UserName: user.UserName,
			CreateAt: time.Now().Unix(),
		},
		AccessToken: token,
	}
	return ctl.RespSuccessWithData(userResp), nil
}
