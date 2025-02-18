package group_share

import (
	"file_service/global"
)

func Create(g *Group) error {
	return global.QY_Db.Create(g).Error
}

func Delete(id uint) error {
	return global.QY_Db.Where("id = ?", id).Delete(&Group{}).Error
}

func GetAllTableData(id uint) (gs []Group, err error) {
	err = global.QY_Db.Preload("Creator").Where("user_id = ?", id).Find(&gs).Error
	return
}
func FindGroupByUID(uid string) (Group, error) {
	var gs Group
	err := global.QY_Db.Where("`key` = ?", uid).First(&gs).Error
	return gs, err
}
func CreateGroupUser(users GroupUsers) error {
	return global.QY_Db.Create(&users).Error
}
func FindGroupUser(val1 uint, val2 uint) (GroupUsers, error) {
	var groupUsers GroupUsers
	err := global.QY_Db.Where("group_id = ? and user_id = ?", val1, val2).First(&groupUsers).Error
	return groupUsers, err
}

func FindGroupUserListByGroupId(val1 map[string]interface{}) ([]FindGroupUsers, error) {
	var groupUsers []FindGroupUsers
	err := global.QY_Db.Table("group_users").
		Select("group_users.user_id AS members_id, users.account_name AS members_name, group_users.created_at AS members_join_time, group_users.group_id, `groups`.label AS group_label").
		Joins("LEFT JOIN users ON users.id = group_users.user_id").
		Joins("LEFT JOIN `groups` ON `groups`.id = group_users.group_id").
		Where(val1).Find(&groupUsers).Error
	return groupUsers, err
}

func FindOrCreateGroupFileListByMap(p *GroupFiles) (GroupFiles, error) {
	err := global.QY_Db.Preload("File").FirstOrCreate(p, &p).Error
	if err != nil {
		return *p, err
	}
	return *p, err
}
func FindGroupFileListByMap(p map[string]interface{}) ([]GroupFiles, error) {
	var groupFiles []GroupFiles
	err := global.QY_Db.Table("group_files").
		Select("group_files.*, users.account_name AS creator_name").
		Joins("left join users ON users.id = group_files.creator_id").
		Where(p).Preload("File").Find(&groupFiles).Error

	return groupFiles, err
}
