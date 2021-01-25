package router

import (
	"diary/controller"
	"diary/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	{
		//登录模块
		engine.GET("/hello", func(context *gin.Context) {
			response := controller.HandleHello()
			context.JSON(200, response)
		})
		engine.POST("/register", func(context *gin.Context) {
			phone := context.PostForm("phone")
			password := context.PostForm("password")
			backDoor := context.PostForm("backDoor")
			fmt.Println("register", phone, password)
			user := model.User{Phone: phone, Password: password}
			response := controller.HandleRegister(&user, backDoor)
			context.JSON(200, response)
		})
		engine.GET("/user/:id", func(context *gin.Context) {
			userid := context.Param("id")
			fmt.Println("get user", userid)
			userId ,_:= strconv.Atoi(userid)
			response := controller.HandleQueryUser(userId)
			context.JSON(200, response)
		})
		engine.DELETE("user/:id", func(context *gin.Context) {
			userid := context.Param("id")
			fmt.Println("delete user", userid)
			userId ,_:= strconv.Atoi(userid)
			response := controller.HandleDeleteUser(userId)
			context.JSON(200, response)
		})
		engine.POST("/login", func(context *gin.Context) {
			phone := context.PostForm("phone")
			password := context.PostForm("password")
			backDoor := context.PostForm("backDoor")
			fmt.Println("login", phone, password)
			user := model.User{Phone: phone, Password: password}
			response := controller.HandleLogin(&user, backDoor)
			context.JSON(200, response)
		})
		engine.GET("/userList", func(context *gin.Context) {
			response := controller.HandleQueryUserList()
			context.JSON(200, response)
		})
	}
	return engine

}
