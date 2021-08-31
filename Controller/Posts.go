package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//增加帖子
func AddPosts(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.AddPosts(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除帖子
func DeletePosts(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeletePosts(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//更新帖子信息
func UpdatePosts(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.UpdatePosts(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取单个帖子信息
func GetPosts(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetPosts(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取帖子列表
func GetPostsList(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetPostsList(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取一个分类下所有帖子
func GetPostsByCategoryId(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetPostsByCategoryId(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//给帖子点赞
func LikePosts(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.LikePosts(c)
	if code == emsg.Success {
		res.Success("", emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//通过标题搜索帖子
func GetPostsListByTitle(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetPostsListByTitle(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
