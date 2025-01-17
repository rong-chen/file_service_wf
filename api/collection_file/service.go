package collection_file

import (
	"errors"
	"file_service/global"
	"gorm.io/gorm"
)

func CreateLikeFile(file LikeFile) (err error) {
	var lf LikeFile
	if err = global.QY_Db.Where("file_id = ? and user_id=?", file.FileId, file.UserId).First(&lf).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没找到，创建
			err = global.QY_Db.Create(&file).Error
		}
		return
	}
	// 如果记录已存在且被软删除，更新软删除标志
	err = global.QY_Db.Unscoped().Where("file_id = ? and user_id = ?", file.FileId, file.UserId).Delete(&LikeFile{}).Error
	return
}

func FindAllListByUserId(userId uint) ([]LikeFile, error) {
	var files []LikeFile
	err := global.QY_Db.Where("user_id = ?", userId).Preload("File").Find(&files).Error
	return files, err
}