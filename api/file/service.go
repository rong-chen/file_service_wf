package file

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"file_service/global"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

func CreateOrFindFile(file File) (resFile File, err error) {
	if errors.Is(global.QY_Db.Where("file_name = ? and file_state = ? and user_id = ? and file_md5 = ? ", file.FileName, true, file.UserId, file.FileMd5).First(&resFile).Error, gorm.ErrRecordNotFound) {
		err = global.QY_Db.Where("file_name = ? and file_md5 = ? and user_id = ?", file.FileName, file.FileMd5, file.UserId).First(&resFile).Error
		// 如果没有找到记录，则根据file模板创建新记录
		file.FileState = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resFile = file                         // 这里将file的内容赋给resFile
			err = global.QY_Db.Create(&file).Error // 创建新记录
			if err != nil {
				return resFile, err
			}
		}
		if err == nil {
			global.QY_Db.Preload("ChunkList").First(&resFile)
		}
	}
	return resFile, err
}

func FindFileRow(userId uint, params QueryParams) (list []File, err error, total int64) {
	query := global.QY_Db.Model(&File{}).Where("user_id = ?", userId)
	if params.isSort {
		query = query.Order("weight desc")
	}
	if params.FileType != "" {
		query = query.Where("file_type = ?  ", params.FileType)
	}
	if params.FileName != "" {
		query = query.Where("file_name LIKE ?", "%"+params.FileName+"%")
	}
	if params.Id != 0 {
		query = query.Where("id = ?", params.Id)
	}
	err = query.Count(&total).Error // 先统计总数
	offset := (params.Page - 1) * params.PageSize
	query = query.Offset(offset).Limit(params.PageSize)
	if err != nil {
		return
	}
	err = query.Preload("ChunkList").Find(&list).Error
	return
}

func CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	// 检查是否已存在相同的 FileId 和 ChunkNumber
	var existingChunk FileChunk
	if err := global.QY_Db.Where("file_id = ? AND chunk_number = ?", id, fileChunkNumber).First(&existingChunk).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var chunk FileChunk
			chunk.ChunkPath = fileChunkPath
			chunk.FileId = id
			chunk.ChunkNumber = fileChunkNumber
			err = global.QY_Db.Create(&chunk).Error
		}
		return err

	}
	return nil
}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func BreakPointContinue(content []byte, fileName string, contentNumber int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber)
	return pathC, err
}

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("文件名或路径不合法")
	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := os.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := os.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil
}

func UpdateFileState(id uint, filePath string) {
	global.QY_Db.Model(&File{}).Where("id = ?", id).Updates(map[string]interface{}{
		"file_state": true,
		"file_path":  filePath,
	})
}

func CollectionFile(i int, id uint) error {
	err := global.QY_Db.Model(&File{}).Where("id = ?", id).Updates(map[string]interface{}{
		"weight": i,
	}).Error
	return err
}

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}

func FindFileById(id uint) (file File, err error) {
	err = global.QY_Db.Model(&File{}).Where("id = ?", id).First(&file).Error
	return
}
