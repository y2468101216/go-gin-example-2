package jwt

import (
	"fmt"
	"gogin/example/pkg/e"
	"gogin/example/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        var data interface{}

        code = e.SUCCESS
        token := c.GetHeader("auth")
		fmt.Println("middleware")
        if token == "" {
            code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
            } else if time.Now().Unix() > claims["expired_at"].(int64) {
                code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }

        if code != e.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : e.GetMsg(code),
                "data" : data,
            })

            c.Abort()
            return
        }

        c.Next()
    }
}