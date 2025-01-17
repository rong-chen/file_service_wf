package user_authority

import "file_service/global"

type UserAuthority struct {
	UserId      uint `json:"user_id" gorm:"column:user_id;comment:用户的id"`
	AuthorityId uint `json:"authority_id" gorm:"column:authority_id;comment:用户权限id"`
	global.QyModel
}

func (u *UserAuthority) TableName() string {
	return "user_authority"
}
