package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r *routes) postRoutes(rg *gin.RouterGroup, mds md.Middleware, post controllers.PostController) {
	rg.Group("/post").
		Use(mds.SetContexts).
		Use(mds.Authentication).
		// POST("/", mds.SetMaxBody, mds.CheckFileLength(4, "files[]"), post.CreatePost).
		POST("/bulk", post.BulkCreatePost).
		DELETE("/:postId", post.DeletePost)
}
