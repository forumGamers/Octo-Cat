package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) commentRoutes(rg *gin.RouterGroup, mds md.Middleware, comment controllers.CommentController) {
	uri := rg.Group("/comment")

	uri.Use(mds.SetContexts)
	uri.Use(mds.Authentication)
	uri.POST("/bulk", comment.BulkComment)
	uri.POST("/:postId", comment.CreateComment)
	uri.DELETE("/:commentId", comment.DeleteComment)
}
