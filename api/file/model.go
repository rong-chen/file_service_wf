package file

import (
	"file_service/global"
)

const (
	breakpointDir = "./breakpointDir/"
	finishDir     = "./uploads/file/fileDir/"
)

type File struct {
	global.QyModel
	FileTotal int         `json:"file_total" gorm:"column:file_total;comment:切片总数" binding:"required"`
	UserId    uint        `json:"user_id" gorm:"column:user_id;comment:上传用户"`
	FileName  string      `json:"file_name" gorm:"column:file_name;comment:文件名称" binding:"required"`
	FilePath  string      `json:"file_path" gorm:"column:file_path;comment:文件路径"`
	FileType  string      `json:"file_type" gorm:"column:file_type;comment:文件格式" binding:"required"`
	FileState bool        `json:"file_state" gorm:"column:file_state;comment:文件状态，是否完成"`
	FileMd5   string      `json:"file_md5" gorm:"column:file_md5;comment:文件md5" binding:"required"`
	ChunkList []FileChunk `json:"chunk_list" gorm:"foreignKey:FileId;references:ID"`
}

type FileChunk struct {
	global.QyModel
	FileId      uint   `json:"file_id" gorm:"column:file_id;comment:文件id"`
	ChunkNumber int    `json:"chunk_number" gorm:"column:chunk_number;comment:文件切片数量"`
	ChunkPath   string `json:"chunk_path" gorm:"column:chunk_path;comment:文件切片路径"`
}

//
//type ExaFileUploadAndDownload struct {
//	global.QyModel
//}
