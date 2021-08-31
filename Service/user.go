package Service

import (
	"Exe/Model"
	ut "Exe/Utils"
	emsg "Exe/Utils/ErrorMessage"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) (*gin.H, int) {
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	if user, code := Model.Register(user.UserName, user.PassWord); code != emsg.Success {
		return nil, code
	} else if actoken, err := ut.GenerateAccessToken(user.UserName); err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else if rftoken, err := ut.GenerateRefreshToken(user.UserName); err != nil {
		return nil, emsg.GenerateRefreshTokenFailed
	} else {
		res := &gin.H{
			"user":          user,
			"access-token":  actoken,
			"refresh-token": rftoken,
		}
		return res, code
	}
}

func Login(c *gin.Context) (*gin.H, int) {
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		return nil, emsg.Error
	}
	if code := Model.Login(user.UserName, user.PassWord); code != emsg.Success {
		return nil, code
	} else if actoken, err := ut.GenerateAccessToken(user.UserName); err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else if rftoken, err := ut.GenerateRefreshToken(user.UserName); err != nil {
		return nil, emsg.GenerateRefreshTokenFailed
	} else {
		res := &gin.H{
			"access-token":  actoken,
			"refresh-token": rftoken,
		}
		return res, code
	}
}
