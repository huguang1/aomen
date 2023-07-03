package routers

import (
	"../apps"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {
	router := gin.Default()  // 使用gin框架来生成路由器
	router.Static("/static", "./static")  // 设置静态文件的访问
	config := router.Group("/config")  // 注册一个路由团体
	{
		config.POST("/login", apps.Login)
		config.GET("/user", apps.UserList)
		config.POST("/adduser", apps.AddUser)
		config.POST("/updateuser", apps.UpdateUser)
		config.POST("/deleteuser", apps.DeleteUser)
		config.GET("/grade", apps.GradeList)
		config.POST("/addgrade", apps.AddGrade)
		config.POST("/updategrade", apps.UpdateGrade)
		config.POST("/deletegrade", apps.DeleteGrade)
		config.GET("/member", apps.ListMember)
		config.POST("/addmember", apps.AddMember)
		config.POST("/updatemember", apps.UpdateMember)
		config.POST("/deletemember", apps.DeleteMember)
		config.GET("/record", apps.RecordList)
		config.POST("/addrecord", apps.AddRecord)
		config.POST("/updaterecord", apps.UpdateRecord)
		config.POST("/deleterecord", apps.DeleteRecord)
	}
	return router
}
