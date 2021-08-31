package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	rep "Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.RefreshToken(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
