package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) replyRoutes(rg *gin.RouterGroup, mds md.Middleware, reply controllers.ReplyController) {
	uri := rg.Group("/reply")

	uri.Use(mds.SetContexts)
	uri.Use(mds.Authentication)
	uri.POST("/:commentId", reply.AddReply)
	uri.DELETE("/:commentId/:replyId", reply.DeleteReply)
}
