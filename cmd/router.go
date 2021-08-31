package cmd

import (
	"Exe/Controller"
	ut "Exe/Utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//用户相关
	r.POST("/register", Controller.Register)
	r.GET("/login", Controller.Login)
	r.GET("/token", ut.RefreshTokenAuth(), Controller.RefreshToken)
	r1 := r.Group("/")
	r1.Use(ut.AccessTokenAuth())
	{
		//分类相关接口
		r1.GET("category/:id", Controller.GetCategory)
		r1.GET("category_list", Controller.GetCategoryList)
		r1.PUT("category/:id", Controller.UpdateCategory)
		r1.POST("category", Controller.AddCategory)
		r1.DELETE("category/:id", Controller.DeleteCategory)
		//帖子相关接口
		r1.GET("posts/:id", Controller.GetPosts)
		r1.POST("posts", Controller.AddPosts)
		r1.DELETE("posts/:id", Controller.DeletePosts)
		r1.GET("posts_list", Controller.GetPostsList)
		r1.PUT("posts/:id", Controller.UpdatePosts)
		r1.GET("posts_list/category/:cateid", Controller.GetPostsByCategoryId)
		r1.PUT("posts/likes/:id", Controller.LikePosts)
		r1.GET("posts/title", Controller.GetPostsListByTitle)
	}
	_ = r.Run(":8080")
}
