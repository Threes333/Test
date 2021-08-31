package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	data, code := Service.Register(c)
	res := rep.NewResponseController(c)
	//判断是否能正确注册
	if code == emsg.Success {
		res.Success(*data, code)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

func Login(c *gin.Context) {
	data, code := Service.Login(c)
	res := rep.NewResponseController(c)
	//判断是否能正确登录
	if code == emsg.Success {
		res.Success(*data, code)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
