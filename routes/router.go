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
	r := routes{gin.Default()}
	r.router.Use(gin.Recovery())

	r.router.Use(mds.CheckOrigin)
	r.router.Use(mds.Cors())
	r.router.Use(logger.SetLogger())

	r.router.MaxMultipartMemory = 100 << 20
	groupRoutes := r.router.Group("/api/v1")
	r.router.POST("/api/v1/post", mds.SetMaxBody, mds.CheckFileLength(4, "files[]"), postController.CreatePost)
	r.postRoutes(groupRoutes, mds, postController)
	r.likeRoutes(groupRoutes, mds, likeController)
	r.commentRoutes(groupRoutes, mds, commentController)
	r.replyRoutes(groupRoutes, mds, replyController)
	r.bookmarkRoutes(groupRoutes, mds, bookmarkController)

	return r.router.Run
}
