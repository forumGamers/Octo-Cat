package controllers

import (
	"github.com/forumGamers/Octo-Cat/pkg/bookmark"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/like"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	p "github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/pkg/reply"
	"github.com/forumGamers/Octo-Cat/pkg/share"
	tp "github.com/forumGamers/Octo-Cat/third-party"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostController interface {
	CreatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	BulkCreatePost(c *gin.Context)
}

type PostControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	Service     p.PostService
	Repo        p.PostRepo
	CommentRepo comment.CommentRepo
	LikeRepo    like.LikeRepo
	ShareRepo   share.ShareRepo
	Ik          tp.ImagekitService
	Validator   *validator.Validate
}

type ReplyController interface {
	AddReply(c *gin.Context)
	DeleteReply(c *gin.Context)
}

type ReplyControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	Service     reply.ReplyService
	CommentRepo comment.CommentRepo
	Validator   *validator.Validate
}

type LikeController interface {
	LikePost(c *gin.Context)
	UnlikePost(c *gin.Context)
	BulkLikes(c *gin.Context)
}

type LikeControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	Service   like.LikeService
	Repo      like.LikeRepo
	PostRepo  p.PostRepo
	Validator *validator.Validate
}

type CommentController interface {
	CreateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	BulkComment(c *gin.Context)
}

type CommentControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	Repo      comment.CommentRepo
	Service   comment.CommentService
	PostRepo  p.PostRepo
	Validator *validator.Validate
}

type BookmarkController interface {
	CreateBookmark(c *gin.Context)
	DeleteBookmark(c *gin.Context)
}

type BookmarkControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	Repo      bookmark.BookmarkRepo
	Service   bookmark.BookmarkService
	PostRepo  post.PostRepo
	Validator *validator.Validate
}
