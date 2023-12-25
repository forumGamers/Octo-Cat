package routes

import (
	"github.com/forumGamers/Octo-Cat/controllers"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes(
	mds md.Middleware,
	postController controllers.PostController,
	likeController controllers.LikeController,
	commentController controllers.CommentController,
	replyController controllers.ReplyController,
	bookmarkController controllers.BookmarkController,
) func(adrs ...string) error {
	r := routes{router: gin.Default()}

	r.router.Use(mds.CheckOrigin)
	r.router.Use(mds.Cors())
	r.router.Use(logger.SetLogger())

	groupRoutes := r.router.Group("/api/v1")
	r.postRoutes(groupRoutes, mds, postController)
	r.likeRoutes(groupRoutes, mds, likeController)
	r.commentRoutes(groupRoutes, mds, commentController)
	r.replyRoutes(groupRoutes, mds, replyController)
	r.bookmarkRoutes(groupRoutes, mds, bookmarkController)

	return r.router.Run
}