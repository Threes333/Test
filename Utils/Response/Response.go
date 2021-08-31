package rep

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回的数据对象
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseController struct {
	*gin.Context
}

func NewResponseController(c *gin.Context) *ResponseController {
	return &ResponseController{c}
}

//成功处理返回对象
func (c *ResponseController) Success(data interface{}, msgcode int) {
	c.JSON(http.StatusOK, &Response{
		Data: data,
		Code: msgcode,
		Msg:  "Success",
	})
}

//处理失败返回信息
func (c *ResponseController) Fail(code int, msg string, msgcode int) {
	c.JSON(code, &Response{
		Code: msgcode,
		Msg:  msg,
		Data: nil,
	})
}
