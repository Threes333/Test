package Service

import (
	ut "Exe/Utils"
	emsg "Exe/Utils/ErrorMessage"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) (*gin.H, int) {
	username := c.PostForm("username")
	actoken, err := ut.GenerateAccessToken(username)
	if err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else {
		return &gin.H{
			"access_token": actoken,
		}, emsg.Success
	}
}
