package apis

import (
	. "../../../app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(context *gin.Context) {
	context.String(http.StatusOK, "It Works!")
}

func GetAll(context *gin.Context) {
	var u User
	users, err := u.GetUserInfo()
	if err != nil {
		log.Fatalln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"data": users,
		"msg":  "success",
	})
}

func GetUserById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	u := User{Id: id}
	_, _ = u.GetUserInfoById()
	if u.Id > 0 {
		context.JSON(http.StatusOK, gin.H{
			"data": u,
			"msg":  "success",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "user not found",
		})
	}
}

func AddUser(context *gin.Context) {
	username := context.Request.FormValue("username")
	passwd := context.Request.FormValue("passwd")
	u := User{UserName: username, PassWd: passwd}
	rs, err := u.AddUser()
	if err != nil {
		log.Fatalln(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": rs,
		"msg":  "Insert success",
	})
}

func UpdateUser(context *gin.Context) {
	username := context.Request.FormValue("username")
	passwd := context.Request.FormValue("passwd")
	id, _ := strconv.Atoi(context.Request.FormValue("id"))
	u := User{Id: id, UserName: username, PassWd: passwd}
	flag, err := u.UpdateUser()
	if flag != true {
		log.Fatalln(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": flag,
		"msg":  "Update success",
	})
}
