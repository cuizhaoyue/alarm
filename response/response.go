package response

import (
	"net/http"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/consts"
	"github.com/gin-gonic/gin"
)

// Response 封装返回的代码
func Response(c *gin.Context, httpStatus, code int, msg string, data interface{}) {
	c.JSON(httpStatus, gin.H{"Code": code, "Message": msg, "Data": data})
}

// Success 成功时的响应
func Success(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, consts.Success, "success", data)
}

// Failed 失败时的响应
func Failed(c *gin.Context, httpStatus, code int, msg string) {
	Response(c, httpStatus, code, msg, nil)
}
