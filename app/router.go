package main

import (
	. "../app/controllers/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.GET("/all", GetAll)
	return router
}
