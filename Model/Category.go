package Model

import (
	"Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	Id   int    `gorm:"type:int;primary_key;auto_increment" json:"id" "`
	Name string `gorm:"type:varchar(255)" json:"name" "`
}

//查询分裂是否存在
func HasCategory(id int) bool {
	var has Category
	if err := DB.Debug().Model(Category{}).Where("id = ?", id).First(&has).Error; err == gorm.ErrRecordNotFound {
		log.Println(err)
		return false
	} else if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//增加分类
func AddCategory(category *Category) int {
	if err := DB.Create(category).Error; err != nil {
		log.Println(err)
		return emsg.CategoryExist
	}
	return emsg.Success
}

//删除分类
func DeleteCategory(id int) int {
	if err := DB.Where("id = ?", id).Delete(Category{}).Error; err != nil {
		log.Println(err)
		return emsg.DeleteCategoryFailed
	}
	return emsg.Success
}

//更改分类
func UpdateCategory(id int, catgory *Category) int {
	var msg = map[string]interface{}{
		"name": catgory.Name,
	}
	if err := DB.Model(Category{}).Where("id = ?", id).Updates(&msg).Error; err != nil {
		return emsg.UpdateCategoryFailed
	}
	return emsg.Success
}

//获取分类信息
func GetCategory(id int) (*Category, int) {
	var category Category
	DB.Debug().Where("id = ?", id).First(&category)
	return &category, emsg.Success
}

//获取分类列表
func GetCategoryList(page, size int) ([]Category, int) {
	var categorys []Category
	if err := DB.Limit(size).Offset((page - 1) * size).Find(&categorys).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, emsg.GetCategoryListFailed
	}
	return categorys, emsg.Success
}
