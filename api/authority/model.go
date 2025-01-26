package authority

import "file_service/global"

type Authorities struct {
	global.QyModel
	AuthorityName string `json:"authority_name" gorm:"column:authority_name;comment:权限名称"`
	BackRouter    string `json:"back_router" gorm:"column:back_router;comment:跳转的路径"`
	AuthorityId   uint   `json:"authority_id" gorm:"column:authority_id;comment:权限ID"`
}

type AuthoritiesMenu struct {
	MenuId      uint `json:"menu_id" gorm:"column:menu_id;comment:菜单ID;primarykey"`
	AuthorityId uint `json:"authority_id" gorm:"column:authority_id;comment:权限ID;primarykey"`
}
