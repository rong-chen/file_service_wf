package file_v2

import (
	"errors"
	"file_service/global"
	"gorm.io/gorm"
)

func CreateOrFindFileInfo(file FileInfo) (FileInfo, error) {
	var files FileInfo
	err := global.QY_Db.Where("file_md5 = ? and user_id = ? ", file.FileMd5, file.UserId).Find(&files).Error
	if err == nil {
		if file.FileName != files.FileName {
			// 如果名称不同，则使用files生成一个
			files.FileName = file.FileName
			global.QY_Db.Create(&files)
		}
		return files, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		global.QY_Db.Create(&file)
		return file, nil
	} else {
		return files, err
	}
}
