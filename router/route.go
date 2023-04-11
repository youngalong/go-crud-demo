package router

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/controller"
)

func SetRoute() *gin.Engine {
	r := gin.Default()
	//增
	r.POST("/user", controller.AddUser)
	//删
	r.DELETE("/user/:id", controller.DeleteUser)
	//改
	r.PUT("/user/:id", controller.UpdateUser)
	//查
	r.GET("/user", controller.GetUserList)
	r.GET("/user/:name", controller.GetUserByName)
	return r
}
