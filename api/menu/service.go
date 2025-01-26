package menu

import (
	"file_service/global"
)

func FindMenuList(menuIds []uint) (menuTree []BaseMenu, err error) {
	var list []BaseMenu
	err = global.QY_Db.Where("id in ?", menuIds).Find(&list).Error
	menuTree = buildMenuTree(list, 0)
	return
}

func buildMenuTree(list []BaseMenu, parentId uint) (menuTree []BaseMenu) {
	var tree []BaseMenu
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
