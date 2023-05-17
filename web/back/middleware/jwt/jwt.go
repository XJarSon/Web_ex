package jwt

import (
	"back/models"
	"back/pkg/e"
	"back/pkg/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT(level int) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.GetHeader("token")

		//token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				msg := models.GetUser(claims.Username)
				if msg.Level > level {
					code = e.ERROR_INSUFFICIENT_ACCESS_RIGHTS
				}
			}
		}
		fmt.Println(code)
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}
		fmt.Println(code)
		c.Next()
	}
}
