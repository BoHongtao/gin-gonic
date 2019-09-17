package apis

import (
	. "../../../app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(context *gin.Context) {
	context.String(http.StatusOK, "It Works!")
}

func GetAll(context *gin.Context)  {
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
