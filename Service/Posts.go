package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/SensitiveWordFilter"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//创建帖子
func AddPosts(c *gin.Context) int {
	var posts Model.Posts
	if err := c.ShouldBindJSON(&posts); err != nil {
		log.Println(err)
		return emsg.Error
	}
	//过滤敏感词
	posts.Title, _ = SensitiveWordFilter.T.Check(posts.Title)
	posts.Content, _ = SensitiveWordFilter.T.Check(posts.Content)
	posts.Describe, _ = SensitiveWordFilter.T.Check(posts.Describe)
	return Model.AddPosts(&posts)
}

//删除帖子
func DeletePosts(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	return Model.DeletePosts(id)
}

//获取帖子信息
func GetPosts(c *gin.Context) (*Model.Posts, int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	return Model.GetPosts(id)
}

//获取帖子列表
func GetPostsList(c *gin.Context) ([]Model.Posts, int) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = -1
	}
	if size <= 0 {
		size = -1
	}
	return Model.GetPostsList(page, size)
}

//更改帖子信息
func UpdatePosts(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	var posts Model.Posts
	if err := c.ShouldBind(&posts); err != nil {
		return emsg.Error
	}
	//过滤敏感词
	posts.Title, _ = SensitiveWordFilter.T.Check(posts.Title)
	posts.Content, _ = SensitiveWordFilter.T.Check(posts.Content)
	posts.Describe, _ = SensitiveWordFilter.T.Check(posts.Describe)
	return Model.UpdatePosts(id, &posts)
}

//获取一个分类下所有帖子
func GetPostsByCategoryId(c *gin.Context) ([]Model.Posts, int) {
	cateid, err := strconv.Atoi(c.Param("cateid"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = -1
	}
	if size <= 0 {
		size = -1
	}
	return Model.GetPostsListByCategoryId(cateid, page, size)
}

//给帖子点赞
func LikePosts(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	return Model.LikePosts(id)
}

//通过标题搜索帖子
func GetPostsListByTitle(c *gin.Context) ([]Model.Posts, int) {
	title := c.PostForm("title")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = -1
	}
	if size <= 0 {
		size = -1
	}
	return Model.GetPostsListByTitle(title, page, size)
}
