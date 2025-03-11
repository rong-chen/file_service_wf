package file_v2

import "file_service/global"

type FileInfo struct {
	FileName       string `json:"file_name" gorm:"column:file_name;comment:文件名称" binding:"required"`
	FileSize       int64  `json:"file_size" gorm:"column:file_size;comment:文件尺寸" binding:"required"`
	FileMd5        string `json:"file_md5" gorm:"column:file_md5;comment:文件md5" binding:"required"`
	FilePath       string `json:"file_path" gorm:"column:file_path;comment:服务器中文件路径" binding:"required"`
	FileSuffix     string `json:"file_suffix" gorm:"column:file_suffix;comment:文件后缀" `
	FileChunkTotal uint   `json:"file_chunk_total" gorm:"column:file_chunk_total;comment:总切片" binding:"required"`
	OverFileTotal  uint   `json:"-" gorm:"column:over_file_total;comment:已完成的切片"`
	UserId         uint   `json:"-" gorm:"column:user_id;comment:上传用户"`
	global.QyModel
}

func (*FileInfo) TableName() string {
	return "file_info"
}
