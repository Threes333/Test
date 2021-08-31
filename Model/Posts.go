package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Posts struct {
	gorm.Model
	Category   Category `gorm:"foreignkey:CateId"`
	UserId     int      `gorm:"type:int" json:"user_id"`                        //帖子作者id
	Username   string   `gorm:"type:varchar(25)" json:"username"`               //帖子作者用户名
	Title      string   `gorm:"type:varchar(50);not null" json:"title"`         //帖子标题
	CateId     int      `gorm:"type:int;not null" json:"cateid"`                //所属分类id
	Describe   string   `gorm:"type:varchar(200)" json:"describe"`              //帖子描述
	Content    string   `gorm:"type:longtext" json:"content"`                   //帖子内容
	LikesCount int      `gorm:"type:int;not null;default:0" json:"likes_count"` //点赞数
	ReadCount  int      `gorm:"type:int;not null;default:0" json:"read_count"`  //阅读量
}

//增加贴子
func AddPosts(posts *Posts) int {
	if err := DB.Create(posts).Error; err != nil {
		log.Println(err)
		return emsg.CreatePostsFailed
	}
	return emsg.Success
}

//删除帖子
func DeletePosts(id int) int {
	if err := DB.Where("id = ?", id).Delete(&Posts{}).Error; err != nil {
		log.Println(err)
		return emsg.DeletePostsFailed
	}
	return emsg.Success
}

//更改帖子
func UpdatePosts(id int, posts *Posts) int {
	var msg = map[string]interface{}{
		"Title":    posts.Title,
		"CateId":   posts.CateId,
		"Describe": posts.Describe,
		"Content":  posts.Content,
	}
	if err := DB.Model(Posts{}).Where("id = ?", id).Updates(&msg).Error; err != nil {
		return emsg.UpdatePostsFailed
	}
	return emsg.Success
}

//获取帖子信息
func GetPosts(id int) (*Posts, int) {
	var posts Posts
	if err := DB.Where("id = ?", id).Preload("Category").First(&posts).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.PostsNoExist
	} else if err != nil {
		return nil, emsg.GetPostsFailed
	}
	DB.Model(Posts{}).Where("id = ?", id).Update("read_count", gorm.Expr("read_count + 1"))
	return &posts, emsg.Success
}

//获取帖子列表
func GetPostsList(page, size int) ([]Posts, int) {
	var posts []Posts
	if err := DB.Preload("Category").Limit(size).Offset((page - 1) * size).Find(&posts).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.PostsNoExist
	} else if err != nil {
		log.Println(err)
		return nil, emsg.GetPostsListFailed
	}
	return posts, emsg.Success
}

//获取一个分类下所有帖子
func GetPostsListByCategoryId(cateid, page, size int) ([]Posts, int) {
	var posts []Posts
	if err := DB.Where("cate_id = ?", cateid).Preload("Category").Limit(size).Offset((page - 1) * size).Find(&posts).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.PostsNoExist
	} else if err != nil {
		return nil, emsg.GetPostsListFailed
	}
	return posts, emsg.Success
}

//给帖子点赞
func LikePosts(id int) int {
	if err := DB.Model(Posts{}).Where("id = ?", id).Update("likes_count", gorm.Expr("likes_count + 1")).Error; err == gorm.ErrRecordNotFound {
		return emsg.PostsNoExist
	} else if err != nil {
		return emsg.LikePostsFailed
	}
	return emsg.Success
}

//通过标题搜索帖子
func GetPostsListByTitle(title string, page, size int) ([]Posts, int) {
	var posts []Posts
	if err := DB.Preload("Category").Limit(size).Offset((page-1)*size).Where("title like ?", title+"%").Find(&posts).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.PostsNoExist
	} else if err != nil {
		return nil, emsg.GetPostsListFailed
	}
	log.Println(err)
	return posts, emsg.Success
}
