package commentApp

import "github.com/gin-gonic/gin"

func CommentRouter(r *gin.Engine) {
	blogs := r.Group("/comment")

	blogs.GET("/:id", GetComments)
	blogs.POST("/create", CreateComments)
}
