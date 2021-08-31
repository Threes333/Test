package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//增加分类
func AddCategory(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.AddCategory(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除分类
func DeleteCategory(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeleteCategory(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//更新分类信息
func UpdateCategory(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.UpdateCategory(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取单个分类信息
func GetCategory(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetCategory(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取分类列表
func GetCategoryList(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetCategoryList(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
