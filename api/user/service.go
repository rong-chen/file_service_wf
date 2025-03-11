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

func (us *service) FindAllUserInfo() (u []Users) {
	global.QY_Db.Find(&u)
	return
}

func (us *service) UpdateIsExamine(id uint, b bool, mountPath string) error {
	return global.QY_Db.Model(&Users{}).Where("id = ?", id).Select("is_examine").Updates(map[string]interface{}{
		"is_examine": b,
		"mount_path": mountPath,
	}).Error
}
func (us *service) UpdateUseDiskSize(id uint, size uint64) error {
	return global.QY_Db.Model(&Users{}).Where("id = ?", id).Select("use_disk_size").Updates(map[string]interface{}{
		"use_disk_size": size,
	}).Error
}
