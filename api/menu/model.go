package menu

import "file_service/global"

type SysBaseMenu struct {
	UID       uint          `json:"uid" gorm:"comment:当前路由ID"`
	ParentId  uint          `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path      string        `json:"path" gorm:"comment:路由path"`        // 路由path
	Name      string        `json:"name" gorm:"comment:路由name"`        // 路由name
	Label     string        `json:"label" gorm:"comment:路由label"`      // 路由name
	Component string        `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Children  []SysBaseMenu `json:"children" gorm:"-"`
	global.QyModel
}
