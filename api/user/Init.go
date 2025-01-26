package user

import (
	"errors"
	"file_service/global"
	"file_service/utils"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

func InitUserDbData() error {
	// 创建超级管理员
	var user Users
	if err := global.QY_Db.Where("id = ?", 1).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.UserName = "超级管理员"
			user.AccountName = "超级管理员"
			user.AuthorityId = 888
			user.UUID = uuid.Must(uuid.NewV4())
			user.Account = "admin"
			user.IsExamine = true
			user.Password = utils.GenerateFromPassword("130561")
			if createErr := global.QY_Db.Create(&user).Error; createErr != nil {
				return fmt.Errorf("创建超级管理员失败: %w", createErr)
			}
		}
	}
	return nil
}
