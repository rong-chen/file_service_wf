package group_share

import (
	"file_service/api/file"
	user2 "file_service/api/user"
	"file_service/global"
	"time"
)

type Group struct {
	Label    string      `json:"label" gorm:"column:label;comment:标题;unique" binding:"required"`
	UserId   uint        `json:"-"  gorm:"column:user_id;comment:创建人"`
	Key      string      `json:"key"  gorm:"column:key;comment:小组的唯一秘钥"`
	Password string      `json:"password" gorm:"column:password;comment:小组的密码"`
	HasPwd   bool        `json:"hasPwd" gorm:"column:has_pwd;comment:是否含有密码"`
	Creator  user2.Users `json:"creator" gorm:"foreignKey:ID;references:UserId;comment: 创建者"`
	global.QyModel
}

type GroupStruct struct {
	Label  string `json:"label" gorm:"column:label;comment:标题;unique" binding:"required"`
	UserId uint   `json:"-"  gorm:"column:user_id;comment:创建人"`
	Key    string `json:"key"  gorm:"column:key;comment:小组的唯一秘钥"`
	HasPwd bool   `json:"hasPwd" gorm:"column:has_pwd;comment:是否含有密码"`
	global.QyModel
}

type GroupUsers struct {
	UserId  uint `json:"user_id" gorm:"column:user_id;comment:加入的用户ID;"`
	GroupId uint `json:"group_id" gorm:"column:group_id;comment:小组ID;"`
	Status  int  `json:"status" gorm:"column:status;comment:状态;default:3"`
	//MembersDetail user2.Users `json:"members_detail" gorm:"foreignKey:ID;references:UserId;comment:加入的用户"`
	//GroupDetail   Group       `json:"group_detail" gorm:"foreignKey:ID;references:GroupId"`
	global.QyModel
}

type FindGroupUsers struct {
	MembersId       uint      `json:"members_id" gorm:"column:members_id;comment:加入的成员ID;"`
	MembersName     string    `json:"members_name" gorm:"column:members_name;comment:加入的成员名称;"`
	MembersJoinTime time.Time `json:"members_join_time" gorm:"column:members_join_time;comment:加入的成员名称;"`
	GroupId         string    `json:"group_id" gorm:"column:group_id;comment:小组ID;"`
	GroupLabel      string    `json:"group_label" gorm:"column:group_label;comment:小组标题;"`
}

type GroupFiles struct {
	FileId      uint      `json:"file_id" gorm:"column:file_id;comment:文件ID;"`
	GroupId     uint      `json:"group_id" gorm:"column:group_id;comment:小组ID;"`
	CreatorId   uint      `json:"creator_id" gorm:"column:creator_id;comment:分享人ID;"`
	CreatorName string    `json:"creator_name" gorm:"-"`
	File        file.File `json:"file" gorm:"foreignKey:ID;references:FileId"`
	global.QyModel
}
