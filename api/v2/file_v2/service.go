package file_v2

import (
	"errors"
	"file_service/global"
	"file_service/utils"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func CreateOrFindFileInfo(file FileInfo) (FileInfo, error, uint) {
	var files FileInfo
	err := global.QY_Db.Where("file_md5 = ? and user_id = ? and file_name = ? ", file.FileMd5, file.UserId, file.FileName).Preload("ChunkList").First(&files).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.QY_Db.Create(&file)
			return file, nil, 0
		}
		return FileInfo{}, err, 1
	}
	return files, nil, 2
}
func FindFileInfoById(id uint) (FileInfo, error) {
	var file FileInfo
	err := global.QY_Db.Where("id = ?", id).First(&file).Error
	return file, err
}

func DeleteFileById(id uint) (err error) {
	err = global.QY_Db.Where("id = ?", id).Delete(&FileInfo{}).Error
	return
}

func InsertFileChunk(info ChunkInfo) error {
	var chunk ChunkInfo
	err := global.QY_Db.Where("file_md5 = ? and chunk_md5 = ? and file_id = ?", info.FileMd5, info.ChunkMd5, info.FileId).First(&chunk).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.QY_Db.Create(&info)
			return nil
		}
	}
	return errors.New("该文件已被上传")
}

func UpdateFileInfo(id uint, filePath string) error {
	var file FileInfo
	err := global.QY_Db.Where("id = ?", id).First(&file).Error
	if err != nil {
		return err
	}
	err = global.QY_Db.Model(&file).Updates(FileInfo{
		FilePath: filePath,
		IsOver:   true,
	}).Error
	return err
}

// MakeFile 生成最终文件
func MakeFile(file FileInfo) (string, error) {
	chunkDir := ""
	finishPath := "" // 完毕的路径
	finishFile := "" // 完毕的文件路径+文件名称和后缀

	finishPath = utils.GetOsPath(filepath.Join(file.FilePath))
	chunkDir = utils.GetOsPath(filepath.Join(file.ChunkPath + "/cache/" + file.FileMd5))
	finishFile = utils.GetOsPath(filepath.Join(file.FilePath, file.FileUUIDName+file.FileSuffix)) // 完毕的文件路径+文件名称和后缀

	// 获取文件块路径
	files, err := os.ReadDir(chunkDir)
	if err != nil {
		return finishPath, err
	}

	// 按 chunk 编号排序
	sort.Slice(files, func(i, j int) bool {
		num1, _ := strconv.Atoi(files[i].Name())
		num2, _ := strconv.Atoi(files[i].Name())
		return num1 < num2
	})

	// 创建目标文件
	_ = os.MkdirAll(finishPath, os.ModePerm)
	fd, err := os.OpenFile(finishFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return finishFile, err
	}
	defer fd.Close()

	// 逐块读取并写入
	buf := make([]byte, 4*1024*1024) // 4MB 缓冲
	for _, file := range files {
		partPath := filepath.Join(chunkDir, file.Name())
		in, err := os.Open(partPath)
		if err != nil {
			_ = os.Remove(finishFile)
			return finishFile, err
		}

		// 按块写入，防止 OOM
		for {
			n, err := in.Read(buf)
			if err != nil && err != io.EOF {
				in.Close()
				return finishFile, err
			}
			if n == 0 {
				break
			}
			if _, err := fd.Write(buf[:n]); err != nil {
				in.Close()
				_ = os.Remove(finishFile)
				return finishFile, err
			}
		}
		in.Close()
	}
	return finishFile, nil
}
func SearchFileInfo(q QueryParams) ([]FileInfo, error) {
	var file []FileInfo
	query := global.QY_Db.Model(&FileInfo{})
	if q.Name != "" {
		query = query.Where("file_name like ?", "%"+q.Name+"%")
	}
	query = query.Where("is_over =  ?", q.IsOver)
	err := query.Limit(q.PageSize).Offset((q.Page - 1) * q.PageSize).Preload("ChunkList").Find(&file).Error
	return file, err
}
