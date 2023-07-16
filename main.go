package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "embed"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/middleware"
	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/router"
	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	"github.com/gin-gonic/gin"
)

func init() {
	path, _ := os.Getwd()
	if err := lib.InitModuleYaml(path+"/conf/", []string{"base", "mysql", "redis"}); err != nil {
		log.Fatal(err)
	}
}

//go:embed swagger/swagger.json
var swaggerByte []byte

func main() {
	// 路由注册
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	// 中间件定义
	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestID())
	engine.Use(middleware.Cors())

	// swagger定义
	swaggerContent := make(map[string]interface{})
	json.Unmarshal(swaggerByte, &swaggerContent)
	engine.GET("/alarmv2/swagger.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, swaggerContent)
	})

	// 路由定义
	router.InitRouter(engine)

	_ = engine.Run(":8080")

}
