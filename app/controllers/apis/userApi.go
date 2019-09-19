package apis

import (
	. "../../../app/models"
	"fmt"
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
	id, _ := strconv.Atoi(context.Request.FormValue("id"))
	username := context.Request.FormValue("username")
	passwd := context.Request.FormValue("passwd")
	u := User{Id: id}
	u.GetUserInfoById()
	if u.Id > 0 {
		u.UserName = username
		u.PassWd = passwd
		ra, err := u.UpdateUser()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update success %d", ra)
		context.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "not found",
		})

	}
}

func DelUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Request.FormValue("id"))
	u := User{Id: id}
	u.GetUserInfoById()
	if u.Id > 0 {
		rs, _err := u.DelUser()
		if _err != nil {
			log.Fatalln(_err)
		}
		msg := fmt.Sprintf("delete success %d", rs)
		context.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})

	}
}
