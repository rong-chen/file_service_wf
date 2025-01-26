package authority

import (
	"errors"
	"file_service/global"
	"fmt"
	"gorm.io/gorm"
)

func InitAuthoritiesData() error {
	// 创建权限列表
	authorityList := []Authorities{
		{AuthorityName: "超级管理员", BackRouter: "", AuthorityId: 888},
		{AuthorityName: "普通用户", BackRouter: "", AuthorityId: 88},
		{AuthorityName: "音乐用户", BackRouter: "", AuthorityId: 8},
	}
	for _, item := range authorityList {
		var author Authorities
		if err := global.QY_Db.Where("authority_id = ?", item.AuthorityId).First(&author).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if createErr := global.QY_Db.Create(&item).Error; createErr != nil {
					return fmt.Errorf("插入权限数据失败: %w", createErr)
				}
			} else {
				return fmt.Errorf("检查权限数据失败: %w", err)
			}
		}
	}
	return nil
}
