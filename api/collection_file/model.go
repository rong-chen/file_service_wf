package collection_file

import (
	"file_service/api/file"
	"file_service/global"
)

type LikeFile struct {
	global.QyModel
	FileId  uint      `json:"file_id" gorm:"column:file_id;comment:文件ID" binding:"required"`
	UserId  uint      `json:"user_id"  gorm:"column:user_id;comment:用户名ID"`
	GroupId uint      `json:"group_id"  gorm:"column:group_id;comment:分组ID"`
	File    file.File `json:"-"  gorm:"foreignKey:FileId;references:ID"`
}
