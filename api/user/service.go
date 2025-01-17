package user

import (
	"errors"
	"file_service/global"
	"file_service/utils"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

var ContextUser = new(service)

type service struct {
}

func (us *service) Create(u Users) (Users, error) {
	var users Users
	if !errors.Is(global.QY_Db.Where("account = ?", u.Account).First(&users).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return users, errors.New("帐号已注册")
	}
	u.Password = utils.GenerateFromPassword(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err := global.QY_Db.Create(&u).Error
	return u, err
}

func (us *service) FindUserInfo(key string, val interface{}) (u Users) {
	global.QY_Db.Where(fmt.Sprintf("%s = ?", key), val).First(&u)
	return
}
