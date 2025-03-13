package menu

import (
	"errors"
	"file_service/api/authority"
	"file_service/global"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var RouterList = []BaseMenu{
	{UID: 1, ParentId: 0, Path: "my_files", Name: "my_files", Label: "我的文件", Icon: "wenjianjia-m", Component: "view/home/index.vue",
		Desc: []uint{
			88, 888,
		}},
	{UID: 2, ParentId: 0, Path: "del_file", Name: "del_file", Label: "回收站", Icon: "huishouzhan", Component: "view/del_file/index.vue",
		Desc: []uint{
			88, 888,
		}},
	{UID: 3, ParentId: 0, Path: "user_manager", Name: "user_manager", Label: "用户管理", Icon: "group-fill", Component: "view/userManager/index.vue",
		Desc: []uint{
			888,
		}},
}

func InitMenuDbData() error {
	// 创建路由列表
	for _, router := range RouterList {
		// 检查是否已存在
		var existing BaseMenu
		if err := global.QY_Db.Where("uid = ?", router.UID).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在时插入
				if createErr := global.QY_Db.Create(&router).Error; createErr != nil {
					return fmt.Errorf("插入路由数据失败: %w", createErr)
				}
			} else {
				return fmt.Errorf("检查路由数据失败: %w", err)
			}
		}
	}
	// 给菜单权限添加默认数据
	var am []authority.AuthoritiesMenu
	for i := range RouterList {
		item := RouterList[i]
		for _, v := range item.Desc {
			am = append(am, authority.AuthoritiesMenu{AuthorityId: v, MenuId: item.UID})
		}
	}
	err := global.QY_Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&am).Error
	return err
}
