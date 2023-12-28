package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) postRoutes(rg *gin.RouterGroup, mds md.Middleware, post controllers.PostController) {
	uri := rg.Group("/post")

	uri.Use(mds.SetContexts)
	uri.Use(mds.Authentication)
	uri.POST("/", mds.SetMaxBody, post.CreatePost)
	uri.POST("/bulk", post.BulkCreatePost)
	uri.DELETE("/:postId", post.DeletePost)

}
