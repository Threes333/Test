package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//增加分类
func AddCategory(c *gin.Context) int {
	var cate Model.Category
	if err := c.ShouldBind(&cate); err != nil {
		log.Println(err)
		return emsg.Error
	}
	if Model.HasCategory(cate.Id) {
		return emsg.CategoryExist
	}
	return Model.AddCategory(&cate)
}

//删除分类
func DeleteCategory(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	if !Model.HasCategory(id) {
		return emsg.CategoryNoExist
	}
	return Model.DeleteCategory(id)
}

//获取分类信息
func GetCategory(c *gin.Context) (*Model.Category, int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	return Model.GetCategory(id)
}

//获取分类列表
func GetCategoryList(c *gin.Context) ([]Model.Category, int) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = -1
	}
	if size <= 0 {
		size = -1
	}
	return Model.GetCategoryList(page, size)
}

//更改分类
func UpdateCategory(c *gin.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	var cate Model.Category
	if err := c.ShouldBind(&cate); err != nil {
		return emsg.Error
	}
	if !Model.HasCategory(id) {
		return emsg.CategoryNoExist
	}
	return Model.UpdateCategory(id, &cate)
}
