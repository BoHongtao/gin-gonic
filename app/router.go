package main

import (
	. "../app/controllers/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.GET("/getAllUser", GetAll)
	router.GET("/getById/:id", GetUserById)
	router.POST("/add", AddUser)
	router.POST("/update", UpdateUser)
	router.POST("/del", DelUser)
	return router
}
