package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	maxAge        = 12
	XRequestIDKey = "X-Request-ID"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowCredentials: false,
		MaxAge:           time.Hour * maxAge,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
}

// RequestID 插入'X-Request-ID'到上下文和每个request/response的header中
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(XRequestIDKey)
		if rid == "" {
			// 在header和上下文中设置请求id
			rid := uuid.NewV4().String()
			c.Request.Header.Set(XRequestIDKey, rid)
			c.Set(XRequestIDKey, rid)
		}
		// 在响应header中设置同样的id
		c.Writer.Header().Set(XRequestIDKey, rid)
		c.Next()
	}
}
