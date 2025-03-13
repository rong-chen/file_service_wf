package userManage

type UpdateParams struct {
	Id        uint   `json:"id" binding:"required"`
	IsExamine bool   `json:"isExamine"`
	MountPath string `json:"mount_path"`
	DiskSize  uint64 `json:"disk_size"`
}
