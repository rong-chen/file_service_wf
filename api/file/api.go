package file

import (
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"io"
	"strconv"
	"strings"
)

func FindFile(c *gin.Context) {
	var file File
	err := c.BindJSON(&file)
	id, _ := c.Get("user_id")
	file.UserId = id.(uint)
	if err != nil {
		response.FailWithMessage("参数错误:"+err.Error(), c)
		return
	}
	file.FilePathName, _ = uuid.NewV1()
	findFile, err := CreateOrFindFile(file)
	if err != nil {
		response.FailWithMessage("添加失败："+err.Error(), c)
		return
	}
	params := make(map[string]interface{})
	params["data"] = findFile
	response.OkWithData(params, "查询成功", c)
}

func FindFileList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	params := parseQueryParams(c)
	row, err, count := FindFileRow(userId.(uint), params)
	if err != nil {
		response.FailWithMessage("获取失败："+err.Error(), c)
		return
	}
	response.OkWithData(map[string]interface{}{
		"list":  row,
		"total": count,
	}, "获取成功", c)
}

func Collection(c *gin.Context) {
	var cp CollectionParams
	err := c.BindJSON(&cp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = CollectionFile(cp.Weight, cp.Id)
	if err != nil {
		response.FailWithMessage("网络错误", c)
		return
	}
	response.OkWithMessage("收藏成功", c)
}

func DownLoadFile(c *gin.Context) {
	fileId := c.Param("fileId")
	fileId = strings.TrimPrefix(fileId, "/")
	uintId, err := strconv.ParseUint(fileId, 10, 0)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	file, err := FindFileById(uint(uintId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.CallBackFile(file.FilePath, file.FileName, c)
}

func FindAllFileList(c *gin.Context) {
	//types, _ := c.GetQuery("type")
	//
	//result := strings.Split(types, "|")
	//row, err := FindFileRow(0, result)
	//if err != nil {
	//	response.FailWithMessage("获取失败："+err.Error(), c)
	//	return
	//}
	//response.OkWithData(row, "获取成功", c)
}

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // 可以继续上传
	} else {
		return false // 切片不完整，废弃
	}
}

func UploadChunkFile(c *gin.Context) {
	// file临时存放点
	fileMd5 := c.Request.FormValue("fileMd5")
	// 文件名称
	fileName := c.Request.FormValue("fileName")
	// 文件片段Md5
	chunkMd5 := c.Request.FormValue("chunkMd5")
	// 文件类型
	fileType := c.Request.FormValue("fileType")
	// 用户id
	userId, _ := c.Get("user_id")
	// file当前的total
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	// fmt.Println(fmt.Sprintf("info ====== 当前传输：%d", chunkNumber))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage("文件接收失败："+err.Error(), c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		response.FailWithMessage("文件读取失败", c)
		return
	}
	cen, _ := io.ReadAll(f)
	//此处要添加md5校验
	if !CheckMd5(cen, chunkMd5) {
		response.FailWithMessage("检查md5失败", c)
		return
	}

	var files File
	files.FileMd5 = fileMd5
	files.FileName = fileName
	files.FileTotal = chunkTotal

	files.FileType = fileType
	files.UserId = userId.(uint)
	findFile, err := CreateOrFindFile(files)
	if err != nil {
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	fileNamePath := findFile.FilePathName.String()
	pathC, err := BreakPointContinue(cen, fileNamePath, chunkNumber, fileMd5)
	if err != nil {
		response.FailWithMessage("断点续传失败", c)
		return
	}
	if err = CreateFileChunk(findFile.ID, pathC, chunkNumber); err != nil {
		response.FailWithMessage("创建文件记录失败$$$"+err.Error(), c)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

func UploadSuccess(c *gin.Context) {
	// file临时存放点
	fileMd5, ok := c.GetQuery("fileMd5")
	// 文件名称
	fileName, ok2 := c.GetQuery("fileName")
	if !ok2 || !ok {
		response.FailWithMessage("参数不完整", c)
		return
	}
	// 用户id
	userId, _ := c.Get("user_id")
	var files File
	files.FileMd5 = fileMd5
	files.FileName = fileName
	files.UserId = userId.(uint)
	findFile, err := CreateOrFindFile(files)
	if err != nil {
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	fileNamePath := findFile.FilePathName.String()
	filePath, err := MakeFile(fileNamePath, fileMd5)
	if err != nil {
		response.FailWithMessage("创建文件失败："+err.Error(), c)
		return
	}
	// 更新file数据库
	UpdateFileState(findFile.ID, filePath)
	err = RemoveChunk(fileMd5)
	if err != nil {
		response.FailWithMessage("删除缓存失败$$$"+err.Error(), c)
		return
	}
	response.Ok(c)
}
