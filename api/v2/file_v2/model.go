package file_v2

import "file_service/global"

type FileInfo struct {
	FileName       string      `json:"file_name" gorm:"column:file_name;comment:文件名称" binding:"required"`
	FileUUIDName   string      `json:"-" gorm:"column:file_uuid_name;comment:文件唯一名称" `
	FileSize       uint64      `json:"file_size" gorm:"column:file_size;comment:文件尺寸" binding:"required"`
	FileMd5        string      `json:"file_md5" gorm:"column:file_md5;comment:文件md5" binding:"required"`
	FilePath       string      `json:"file_path" gorm:"column:file_path;comment:服务器中文件路径" binding:"required"`
	IsOver         bool        `json:"is_over"  gorm:"column:is_over;comment:是否上传完毕;default:false"`
	ChunkPath      string      `json:"chunk_path" gorm:"column:chunk_path;comment:切片总路径" `
	FileSuffix     string      `json:"file_suffix" gorm:"column:file_suffix;comment:文件后缀" `
	FileChunkTotal uint        `json:"file_chunk_total" gorm:"column:file_chunk_total;comment:总切片" binding:"required"`
	UserId         uint        `json:"-" gorm:"column:user_id;comment:上传用户"`
	ChunkList      []ChunkInfo `json:"chunk_list" gorm:"foreignKey:FileId;references:ID"`
	global.QyModel
}

type ChunkInfo struct {
	FileMd5  string `json:"file_md5" gorm:"column:file_md5;comment:总文件md5" binding:"required"`
	FileId   uint   `json:"file_id" gorm:"column:file_id;comment:总文件id" binding:"required"`
	ChunkMd5 string `json:"chunk_md5" gorm:"chunk_md5:file_md5;comment:切片md5" binding:"required"`
	Index    string `json:"index" gorm:"column:index;comment:当前文件切片的index" binding:"required"`
	SavePath string `json:"-" gorm:"column:save_path;comment:当前文件切片保存路径"`
	global.QyModel
}

type QueryParams struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
	Name     string `form:"name" binding:"omitempty"`
	IsOver   bool   `form:"is_over" binding:"omitempty"`
}

func (*FileInfo) TableName() string {
	return "file_info"
}
