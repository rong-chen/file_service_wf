package authority

import "file_service/global"

func FindAuthorities(id uint) ([]AuthoritiesMenu, error) {
	var am []AuthoritiesMenu
	err := global.QY_Db.Where("authority_id = ?", id).Find(&am).Error
	return am, err
}
