package group_share

import (
	"errors"
	"file_service/model/common/response"
	"file_service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func CreateGroup(c *gin.Context) {
	id, _ := c.Get("user_id")
	type Params struct {
		Label         string `json:"label" binding:"required"`
		IsHasPassword bool   `json:"hasPwd" binding:"omitempty"`
		Password      string `json:"password"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage("参数有误"+err.Error(), c)
		return
	}

	if p.IsHasPassword && p.Password == "" {
		response.FailWithMessage("参数有误", c)
		return
	}
	uid, _ := uuid.NewV1()
	suid := uid.String()
	suid = strings.ReplaceAll(suid, "-", "")
	generateFromPassword := ""
	if p.IsHasPassword {
		generateFromPassword = utils.GenerateFromPassword(p.Password)
	}

	g := &Group{
		Label:    p.Label,
		UserId:   id.(uint),
		Password: generateFromPassword,
		HasPwd:   p.IsHasPassword,
		Key:      suid,
	}
	err = Create(g)
	if err != nil {
		response.FailWithMessage("添加失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("新增成功", c)
}

func DeleteGroup(c *gin.Context) {
	id, _ := c.GetQuery("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	err = Delete(uint(ids))
	if err != nil {
		response.FailWithMessage("网络有误"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func GetGroup(c *gin.Context) {
	id, _ := c.Get("user_id")
	data, err := GetAllTableData(id.(uint))
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(data, "获取成功", c)
}
func Join(c *gin.Context) {
	id, _ := c.Get("user_id")
	type Params struct {
		ID       string `json:"id" binding:"required"`
		Password string `json:"password"`
	}
	var p Params
	err := c.BindJSON(&p)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	share, err := FindGroupByUID(p.ID)
	if err != nil {
		response.FailWithMessage("网络错误"+err.Error(), c)
		return
	}
	if share.UserId == id.(uint) {
		response.FailWithMessage("添加失败，请勿添加个人小组", c)
		return
	}
	if share.Password != "" && !utils.CompareHashAndPassword(share.Password, p.Password) {
		response.FailWithMessage("密码错误", c)
		return
	}
	gUser, err := FindGroupUser(share.ID, id.(uint))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("网络错误"+err.Error(), c)
		return
	}
	if gUser.ID != 0 {
		response.FailWithMessage("请勿重复添加", c)
		return
	}
	var groupUser GroupUsers
	groupUser.GroupId = share.ID
	groupUser.UserId = id.(uint)
	err = CreateGroupUser(groupUser)
	if err != nil {
		response.FailWithMessage("新增失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("新增成功", c)
}

func FindGroupUsersList(c *gin.Context) {
	groupId, is := c.GetQuery("group_id")
	params := map[string]interface{}{}
	if is {
		ids, _ := strconv.Atoi(groupId)
		params["group_users.group_id"] = ids
	}

	userId, isUserId := c.GetQuery("members_id")
	if isUserId {
		params["group_users.user_id"] = userId
	}
	list, err := FindGroupUserListByGroupId(params)
	if err != nil {
		response.FailWithMessage("查询失败"+err.Error(), c)
		return
	}
	response.OkWithData(list, "查询成功", c)
}

func AddFile(c *gin.Context) {
	type Params struct {
		FileId  uint `json:"file_id" binding:"required"`
		GroupId uint `json:"group_id" binding:"required"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	id, _ := c.Get("user_id")
	queryParams := &GroupFiles{
		GroupId:   p.GroupId,
		FileId:    p.FileId,
		CreatorId: id.(uint),
	}
	groupFiles, err := FindOrCreateGroupFileListByMap(queryParams)
	if err != nil {
		response.FailWithMessage("添加失败"+err.Error(), c)
		return
	}
	response.OkWithData(groupFiles, "新增成功", c)
}

func FindGroupFilesList(c *gin.Context) {
	groupId, is := c.GetQuery("group_id")
	params := map[string]interface{}{}
	if is {
		ids, _ := strconv.Atoi(groupId)
		params["group_files.group_id"] = ids
	}
	groupFiles, err := FindGroupFileListByMap(params)
	if err != nil {
		response.FailWithMessage("添加失败"+err.Error(), c)
		return
	}
	response.OkWithData(groupFiles, "获取成功", c)
}
