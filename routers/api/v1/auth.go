package v1

import (
	"gogin/example/models"
	"gogin/example/pkg/e"
	"gogin/example/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	var query models.GetAuthForm
	code := e.INVALID_PARAMS
    if err := c.ShouldBindJSON(&query); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
    }

	if !models.CheckAuth(query.Username, query.Password) {
		code = e.ERROR_AUTH
		c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
	}

	data := make(map[string]interface{})
	data["token"] = util.GenerateToken(query.Username)

	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

func ParseUserToken(c *gin.Context) {
	user, err := util.ParseToken(c.GetHeader("auth"))

	code := e.ERROR_AUTH_CHECK_TOKEN_FAIL
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "code": code,
            "msg": e.GetMsg(code),
            "data": make(map[string]string),
        })
        return
	}

	code = e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : user,
    })
}