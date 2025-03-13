package userManage

import (
	user2 "file_service/api/user"
	"file_service/global"
)

func FindAllUserInfo() (u []user2.Users) {
	global.QY_Db.Find(&u)
	return
}

func UpdateUsers(params UpdateParams) error {
	updateData := map[string]interface{}{}
	updateData["is_examine"] = params.IsExamine
	if params.MountPath != "" {
		updateData["mount"] = params.MountPath
	}
	if params.DiskSize >= 0 {
		updateData["disk_size"] = params.DiskSize
	}
	return global.QY_Db.Model(&user2.Users{}).Where("id = ?", params.Id).Select("is_examine", "mount", "disk_size").Updates(updateData).Error
}
