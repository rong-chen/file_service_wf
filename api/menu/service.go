package menu

import (
	"errors"
	"file_service/global"
	"fmt"
	"gorm.io/gorm"
)

func InitRouterDb() error {
	routerList := []SysBaseMenu{
		{UID: 1, ParentId: 0, Path: "home", Name: "home", Label: "仪表盘", Component: "view/superAdmin/index.vue"},
		{UID: 2, ParentId: 0, Path: "admin", Name: "superAdmin", Label: "超级管理员", Component: "view/superAdmin/index.vue"},
		{UID: 3, ParentId: 0, Path: "file", Name: "file", Label: "文件操作", Component: "view/superAdmin/index.vue"},
		{UID: 4, ParentId: 3, Path: "upload_file", Name: "upload_file", Label: "文件上传", Component: "view/superAdmin/index.vue"},
	}
	for _, router := range routerList {
		// 检查是否已存在
		var existing SysBaseMenu
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
	return nil
}

func FindMenuList() (menuTree []SysBaseMenu, err error) {
	var list []SysBaseMenu
	err = global.QY_Db.Find(&list).Error
	menuTree = buildMenuTree(list, 0)
	return
}

func buildMenuTree(list []SysBaseMenu, parentId uint) (menuTree []SysBaseMenu) {
	var tree []SysBaseMenu
	for _, menu := range list {
		if menu.ParentId == parentId {
			// 查找当前菜单的子菜单
			children := buildMenuTree(list, menu.UID)
			menu.Children = children
			tree = append(tree, menu)
		}
	}
	return tree
}
