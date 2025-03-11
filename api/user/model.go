package user

import (
	"file_service/global"
	"github.com/gofrs/uuid/v5"
)

type Users struct {
	UUID           uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户uuid" `
	UserName       string    `json:"user_name" form:"userName" gorm:"column:user_name;comment:用户姓名"`
	AccountName    string    `json:"account_name" form:"account_name" gorm:"column:account_name;comment:账号名称" binding:"required"`
	Account        string    `json:"account" form:"account"  gorm:"column:account;comment:账号" binding:"required"`
	Password       string    `json:"password" form:"password"  gorm:"column:password;comment:密码" binding:"required"`
	ProfilePicture string    `json:"profile_picture" gorm:"column:profile_picture;comment:头像"`
	AuthorityId    uint      `json:"authority_id" gorm:"column:authority_id;comment:权限ID"`
	IsExamine      bool      `json:"isExamine" gorm:"column:is_examine;comment:是否通过审核;default:false"`
	DiskSize       uint64    `json:"disk_size"  gorm:"column:disk_size;comment:用户可使用磁盘大小;default:0"`
	UseDiskSize    uint64    `json:"use_disk_size"  gorm:"column:use_disk_size;comment:已使用;default:0"`
	MountPath      string    `json:"mount_path"  gorm:"column:mount;comment:挂载目录;"`
	global.QyModel
}

func (*Users) TableName() string {
	return "users"
}
