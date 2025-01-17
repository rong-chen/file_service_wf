package file_secret

import "file_service/global"

type FileSecret struct {
	global.QyModel
	Desc string `json:"desc" gorm:"column:desc;comment:备注"`
	Name string `json:"name" gorm:"column:name;comment:名称" binding:"required"`

	CreateUser uint   `json:"create_user" gorm:"column:create_user;comment:创建者ID" binding:"required"`
	Key        string `json:"key" gorm:"column:key;comment:key" binding:"required"`
	Secret     string `json:"secret" gorm:"column:secret;comment:密钥" binding:"required"`
	FileType   string `json:"file_type" gorm:"column:file_type;comment:可见文件类型" binding:"required"`
	Disabled   bool   `json:"disabled" gorm:"column:disabled;comment:是否禁用" binding:"required"`
}
