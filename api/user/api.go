package user

import (
	"file_service/model/common/requests"
	"file_service/model/common/response"
	"file_service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
)

func Login(c *gin.Context) {
	type Params struct {
		Account  string `form:"account" json:"account" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage("帐号密码有误"+err.Error(), c)
		return
	}
	u := ContextUser.FindUserInfo("account", p.Account)
	if !utils.CompareHashAndPassword(u.Password, p.Password) {
		response.FailWithMessage("帐号密码有误", c)
		return
	}
	if !u.IsExamine {
		requests.NoAuthority("请等待管理员审核", c)
		return
	}
	token, err := utils.JWTAPP.CreateToken(u.ID)
	if err != nil {
		response.FailWithMessage("token创建失败"+err.Error(), c)
		return
	}
	response.OkWithData(map[string]interface{}{
		"token": token,
	}, "登陆成功", c)
}

func RegisterUser(c *gin.Context) {
	var users Users
	err := c.ShouldBind(&users)
	if err != nil {
		// 有必填项未填写
		response.FailWithMessage("请填写完整："+err.Error(), c)
		return
	}
	file, _ := c.FormFile("file")
	if file != nil {
		uid, _ := uuid.NewV1()
		dsn := fmt.Sprintf("/uploads/file/image/profile/%s.png", uid.String())
		users.ProfilePicture = dsn
		err := c.SaveUploadedFile(file, "."+dsn)
		if file.Size > 1024*1024*5 {
			// 返回错误码，不允许大于5M
			response.FailWithMessage("图片尺寸不能大于5M", c)
			return
		}
		if err != nil {
			// 保存错误
			response.FailWithMessage("图头像上传失败"+err.Error(), c)
			return
		}
	}
	users.AuthorityId = 88
	_, err = ContextUser.Create(users)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("注册失败,%s", err.Error()), c)
		return
	}
	// 创建用户成功
	response.OkWithMessage("注册成功", c)
}

func List(c *gin.Context) {
	authorityId, ok := c.Get("authorityId")
	if !ok {
		response.FailWithMessage("用户权限获取错误", c)
		return
	}
	uAuthorityId := authorityId.(uint)
	if uAuthorityId != 888 {
		response.FailWithMessage("暂无权限", c)
		return
	}
	list := ContextUser.FindAllUserInfo()
	response.OkWithData(map[string]interface{}{
		"list": list,
	}, "获取成功", c)
}

func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("user_id")
	u := ContextUser.FindUserInfo("id", userId)
	response.OkWithData(u, "获取成功", c)
}

func ConsentRegister(c *gin.Context) {
	authorityId, ok := c.Get("authorityId")
	if !ok {
		response.FailWithMessage("用户权限获取错误", c)
		return
	}
	uAuthorityId := authorityId.(uint)
	if uAuthorityId != 888 {
		response.FailWithMessage("暂无权限", c)
		return
	}
	type Params struct {
		Id        uint `json:"id" binding:"required"`
		IsExamine bool `json:"isExamine"`
	}

	var p Params
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage("变更失败"+err.Error(), c)
		return

	}
	if p.IsExamine != true && p.IsExamine != false {
		response.FailWithMessage("参数错误", c)
	}

	err = ContextUser.UpdateIsExamine(p.Id, p.IsExamine)
	if err != nil {
		response.FailWithMessage("变更失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("获取成功", c)
}
