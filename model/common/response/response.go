package response

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func CallBackFile(filePath string, fileName string, c *gin.Context) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			FailWithMessage("文件不存在", c)
			return
		} else {
			FailWithMessage("检查文件时发生错误", c)
			return
		}
	} else {
		// 手动打开文件并传输
		file, err := os.Open(filePath)
		if err != nil {
			FailWithMessage("文件打开失败", c)
			return
		}
		defer file.Close()

		// 获取文件大小并设置响应头
		fileInfo, err := file.Stat()
		if err != nil {
			FailWithMessage("获取文件信息失败", c)
			return
		}
		size := strconv.FormatInt(fileInfo.Size(), 10)
		c.Header("Content-Type", "application/octet-stream")
		encodedFileName := url.QueryEscape(fileName)
		c.Header("Content-Disposition", "attachment; filename*=UTF-8''"+encodedFileName)
		c.Header("Content-Length", size) // 明确设置文件大小
		// 将文件流传输给客户端
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			FailWithMessage("文件传输失败", c)
			return
		}
	}
}
