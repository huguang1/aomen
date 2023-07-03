package main

import (
	db "./dbs"
	"./libs"
	"./routers"
	"github.com/gin-gonic/gin"
)

func main() {
	defer db.Conns.Close() // 延后关闭数据库连接
	gin.SetMode(gin.DebugMode)  // 设置项目为debug模式，gin.SetMode(gin.ReleaseMode)生产模式
	router := routers.InitRouter()  // 生成路由器
	router.Run(":" + libs.Conf.Read("site", "httpport"))  // 路由器运行

}
