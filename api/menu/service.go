package menu

import (
	"errors"
	"file_service/api/authority"
	user2 "file_service/api/user"
	"file_service/global"
	"file_service/utils"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

func InitRouterDb() error {

	// 创建路由列表
	routerList := []SysBaseMenu{
		{UID: 1, ParentId: 0, Path: "home", Name: "home", Label: "仪表盘", Component: "view/Home/index.vue"},
		{UID: 2, ParentId: 0, Path: "admin", Name: "superAdmin", Label: "超级管理员", Component: "view/superAdmin/index.vue"},
		{UID: 3, ParentId: 0, Path: "file", Name: "file", Label: "文件操作"},
		{UID: 4, ParentId: 3, Path: "upload_file", Name: "upload_file", Label: "文件上传", Component: "view/file/upload/index.vue"},
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
	// 创建权限列表
	authorityList := []authority.Authorities{
		{AuthorityName: "超级管理员", BackRouter: "", AuthorityId: 888},
		{AuthorityName: "普通用户", BackRouter: "", AuthorityId: 88},
		{AuthorityName: "音乐用户", BackRouter: "", AuthorityId: 8},
	}

	for _, item := range authorityList {
		var author authority.Authorities
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

	// 创建超级管理员
	var user user2.Users
	if err := global.QY_Db.Where("id = ?", 1).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.UserName = "超级管理员"
			user.AccountName = "超级管理员"
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
