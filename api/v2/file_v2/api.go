package file_v2

import (
	"errors"
	"file_service/api/user"
	"file_service/model/common/response"
	"file_service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
)

// CheckFile 该方法创建file，并且检测file是否已经上传了
func CheckFile(c *gin.Context) {
	id, _ := c.Get("user_id")
	var fileInfo FileInfo
	err := c.ShouldBindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage("参数错误:"+err.Error(), c)
		return
	}
	fileInfo.UserId = id.(uint)
	userInfo := user.ContextUser.FindUserInfo("id", id)
	u := userInfo.UseDiskSize + fileInfo.FileSize
	if u > userInfo.DiskSize {
		response.FailWithMessage("磁盘空间不足", c)
		return
	}
	err = user.ContextUser.UpdateUseDiskSize(fileInfo.UserId, u)
	if err != nil {
		response.FailWithMessage("网络错误"+err.Error(), c)
		return
	}
	v1, _ := uuid.NewV1()
	fileInfo.ChunkPath = utils.GetOsPath(filepath.Join(fileInfo.FilePath + "/cache-file/" + v1.String()))
	fileInfo.FileUUIDName = v1.String()
	file, err, n := CreateOrFindFileInfo(fileInfo)
	if n == 2 {
		response.OkWithData(file, "文件已存在", c)
		return
	}
	if err != nil {
		response.FailWithMessage("检测文件错误:"+err.Error(), c)
		return
	}
	response.OkWithData(file, "检测文件成功", c)
}

func UploadChunk(c *gin.Context) {
	index := c.PostForm("index")
	fileMd5 := c.PostForm("file_md5")
	chunkMd5 := c.PostForm("chunk_md5")
	chunkPath := c.PostForm("chunk_path")
	fileId := c.PostForm("file_id")
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("上传切片失败:"+err.Error(), c)
		return
	}

	dst := utils.GetOsPath(filepath.Join(chunkPath+"/cache/"+fileMd5, index))

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		response.FailWithMessage("上传切片失败:"+err.Error(), c)
		return
	}

	iFid, _ := strconv.Atoi(fileId)

	var chunk = ChunkInfo{
		ChunkMd5: chunkMd5,
		Index:    index,
		FileMd5:  fileMd5,
		SavePath: dst,
		FileId:   uint(iFid),
	}
	err = InsertFileChunk(chunk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("上传文件成功", c)
}

// CombinedFile 上传成功
func CombinedFile(c *gin.Context) {
	type Params struct {
		Id uint `json:"id"`
	}
	var params Params
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("参数有误"+err.Error(), c)
		return
	}
	file, err := FindFileInfoById(params.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("暂无文件记录，请重新上传"+err.Error(), c)
			return
		}
		response.FailWithMessage("文件合并失败"+err.Error(), c)
		return
	}

	finishFile, err := MakeFile(file)
	if err != nil {
		response.FailWithMessage("文件合并失败"+err.Error(), c)
		return
	}

	// ✅ 合并成功后，删除碎片目录
	err = os.RemoveAll(utils.GetOsPath(filepath.Join(file.ChunkPath)))
	if err != nil {
		response.FailWithMessage("文件更新失败"+err.Error(), c)
		return
	}

	err = UpdateFileInfo(params.Id, finishFile)
	if err != nil {
		response.FailWithMessage("文件更新失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("文件合并成功", c)
}

// List 获取列表
func List(c *gin.Context) {
	var q QueryParams
	err := c.ShouldBindJSON(&q)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	info, err := SearchFileInfo(q)
	if err != nil {
		response.FailWithMessage("网络错误"+err.Error(), c)
		return
	}
	response.OkWithData(info, "获取成功", c)
	return
}

func Delete(c *gin.Context) {
	id, ok := c.GetQuery("id")
	userId := c.MustGet("user_id").(uint)
	if !ok {
		response.FailWithMessage("网络异常", c)
		return
	}
	uintId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		response.FailWithMessage("参数异常", c)
		return
	}
	file, err := FindFileInfoById(uint(uintId))
	if err != nil {
		response.FailWithMessage("网络异常"+err.Error(), c)
		return
	}
	if file.UserId != userId {
		response.FailWithMessage("无权利删除", c)
		return
	}

	if file.FilePath != "" {
		_, err := os.Stat(file.FilePath)
		if err == nil {
			err = os.Remove(file.FilePath)
			if err != nil {
				response.FailWithMessage("删除文件失败:"+err.Error(), c)
				return
			}
		}
	} else {
		response.FailWithMessage("网络异常", c)
		return
	}

	err = DeleteFileById(file.ID)
	if err != nil {
		response.FailWithMessage("网络异常"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
