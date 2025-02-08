package file

import (
	"file_service/global"
	"github.com/gofrs/uuid/v5"
)

const (
	breakpointDir = "./breakpointDir/"
	finishDir     = "./uploads/file/fileDir/"
)

type File struct {
	global.QyModel
	FileTotal        int         `json:"file_total" gorm:"column:file_total;comment:切片总数" binding:"required"`
	UserId           uint        `json:"user_id" gorm:"column:user_id;comment:上传用户"`
	FileName         string      `json:"file_name" gorm:"column:file_name;comment:文件名称" binding:"required"`
	FilePathName     uuid.UUID   `json:"file_path_name" gorm:"column:file_path_name;comment:文件路径名称"`
	FilePath         string      `json:"file_path" gorm:"column:file_path;comment:文件路径"`
	FileType         string      `json:"file_type" gorm:"column:file_type;comment:文件格式" binding:"required"`
	FileState        bool        `json:"file_state" gorm:"column:file_state;comment:文件状态，是否完成;default:false"`
	FileMd5          string      `json:"file_md5" gorm:"column:file_md5;comment:文件md5" binding:"required"`
	Widget           int         `json:"weight" gorm:"column:weight;comment:文件权重;default:1"`
	FileSize         uint64      `json:"file_size" gorm:"column:file_size;comment:文件大小;default:0" binding:"required"`
	ChunkList        []FileChunk `json:"chunk_list" gorm:"foreignKey:FileId;references:ID"`
	IsShare          bool        `json:"is_share" gorm:"column:is_share;comment:是否是共享文件;default:false"`
	ShareUserId      uint        `json:"share_user_id" gorm:"column:share_user_id;comment:共享用户ID"`
	ShareAccountName string      `json:"share_account_name" gorm:"column:share_account_name;comment:共享用户帐号名称"`
}

// 文件切片
type FileChunk struct {
	global.QyModel
	FileId      uint   `json:"file_id" gorm:"column:file_id;comment:文件id"`
	ChunkNumber int    `json:"chunk_number" gorm:"column:chunk_number;comment:文件切片数量"`
	ChunkPath   string `json:"chunk_path" gorm:"column:chunk_path;comment:文件切片路径"`
	ParsePath   string `json:"parse_path" gorm:"column:parse_path;comment:切片父路径"`
}

// 查询参数
type QueryParams struct {
	FileType string
	FileName string
	Page     int
	PageSize int
	isSort   bool
	Id       uint
}

// 收藏文件
type CollectionParams struct {
	Id     uint `json:"id" binding:"required"`
	Weight int  `json:"weight" binding:"required"`
}

// 共享文件
type ShareFileInfo struct {
	FromUser uint `json:"from_user" gorm:"column:from_user;comment:从哪用户来" binding:"required"`
	ToUser   uint `json:"to_user" gorm:"column:to_user;comment:到哪个用户去" binding:"required"`
	FileId   uint `json:"file_id" gorm:"column:file_id;comment:文件ID" binding:"required"`
	global.QyModel
}
