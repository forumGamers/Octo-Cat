package main

import (
	"os"

	"github.com/forumGamers/Octo-Cat/controllers"
	"github.com/forumGamers/Octo-Cat/errors"
	md "github.com/forumGamers/Octo-Cat/middlewares"
	"github.com/forumGamers/Octo-Cat/pkg/bookmark"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/like"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/pkg/reply"
	"github.com/forumGamers/Octo-Cat/pkg/share"
	"github.com/forumGamers/Octo-Cat/routes"
	tp "github.com/forumGamers/Octo-Cat/third-party"
	v "github.com/forumGamers/Octo-Cat/validations"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/joho/godotenv"
)

func main() {
	errors.PanicIfError(godotenv.Load())

	validate := v.GetValidator()
	ik := tp.NewImageKit()
	w := web.NewResponseWriter()
	r := web.NewRequestReader()
	mds := md.NewMiddlewares(w, r)

	postRepo := post.NewPostRepo()
	likeRepo := like.NewLikeRepo()
	commentRepo := comment.NewCommentRepo()
	shareRepo := share.NewShareRepo()
	bookmarkRepo := bookmark.NewBookMarkRepo()

	postService := post.NewPostService(postRepo)
	likeService := like.NewLikeService(likeRepo)
	commentService := comment.NewCommentService(commentRepo)
	replyService := reply.NewReplyService(commentRepo)
	bookmarkService := bookmark.NewBookMarkService(bookmarkRepo)

	postController := controllers.NewPostControllers(w, r, postService, postRepo, commentRepo, likeRepo, shareRepo, ik, validate)
	likeController := controllers.NewLikeController(w, r, likeService, likeRepo, postRepo, validate)
	commentController := controllers.NewCommentController(w, r, commentRepo, commentService, postRepo, validate)
	replyController := controllers.NewReplyController(w, r, replyService, commentRepo, validate)
	bookmarkController := controllers.NewBookmarkController(w, r, bookmarkRepo, bookmarkService, postRepo, validate)

	app := routes.NewRoutes(mds, postController, likeController, commentController, replyController, bookmarkController)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4300"
	}

	app(":" + port)
}
