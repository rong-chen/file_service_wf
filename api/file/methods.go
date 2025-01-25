package file

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func parseQueryParams(c *gin.Context) QueryParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	return QueryParams{
		FileType: c.Query("fileType"),
		FileName: c.Query("fileName"),
		Page:     page,
		PageSize: pageSize,
		isSort:   c.Query("isSort") == "是",
	}
}
