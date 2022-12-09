package blogApp

import "github.com/gin-gonic/gin"

func BlogRouter(r *gin.Engine) {
	blogs := r.Group("/blog")

	blogs.GET("/:id", GetBlogById)
	blogs.GET("/", GetBlogs)
	blogs.POST("/create", CreateBlog)
	blogs.GET("/delete/:id", DeleteBlog)
	blogs.POST("/update/:id", UpdateBlog)
}
