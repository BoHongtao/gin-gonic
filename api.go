package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	router := gin.Default()
	//get请求
	router.GET("/getall",getall)
	router.GET("/getbyname",getbyname)

	//post请求
	router.POST("postall",postall)


	//上传文件大小限制（默认32M）
	router.MaxMultipartMemory = 8
	//上传单个文件
	router.POST("upload",upload)
	router.POST("uploads",uploads)

	//分组  v1/login
	v1 := router.Group("v1")
	{
		v1.POST("login",login)
		v1.POST("logout",logout)
	}

	router.Run()
}



func login(context *gin.Context) {
	context.JSON(200,gin.H{
		"message":"v1-login",
	})
}

func logout(context *gin.Context) {
	context.JSON(200,gin.H{
		"message":"v1-logout",
	})
}


//get方式
func getall(context *gin.Context) {
	context.JSON(200,gin.H{
		"message":"get all name",
	})
}


//get方式获取参数
func getbyname(context *gin.Context) {
	firstname := context.DefaultQuery("firstname","")
	lastname := context.DefaultQuery("lastname","")
	//context.String(http.StatusOK, "Hello %s", name)
	context.JSON(200,gin.H{
		"firstname":firstname,
		"lastname":lastname,
	})
}

//post方式获取参数
func postall(context *gin.Context) {
	message := context.DefaultPostForm("message","")
	context.JSON(200,gin.H{
		"message":message,
	})
}

//上传单个文件
func upload(context *gin.Context) {
	file,_:=context.FormFile("file")
	log.Println(file.Filename)
	var dst = "./"
	context.SaveUploadedFile(file, dst)
	context.JSON(200,gin.H{
		"message":file.Filename,
	})
}

//上传多个文件
func uploads(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["upload[]"]
	var dst = "./"
	for _, file := range files {
		log.Println(file.Filename)
		context.SaveUploadedFile(file, dst)
	}
	context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
