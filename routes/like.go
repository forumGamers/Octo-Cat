package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r *routes) likeRoutes(rg *gin.RouterGroup, mds md.Middleware, like controllers.LikeController) {
	uri := rg.Group("/like")

	uri.Use(mds.SetContexts)
	uri.Use(mds.Authentication)
	uri.POST("/bulk", like.BulkLikes)
	uri.POST("/:postId", like.LikePost)
	uri.DELETE("/:postId", like.UnlikePost)
}
