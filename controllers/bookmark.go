package controllers

import (
	"context"
	"strings"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/bookmark"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBookmarkController(
	w web.ResponseWriter,
	r web.RequestReader,
	repo bookmark.BookmarkRepo,
	service bookmark.BookmarkService,
	postRepo post.PostRepo,
	validate *validator.Validate,
) BookmarkController {
	return &BookmarkControllerImpl{w, r, repo, service, postRepo, validate}
}

func (bc *BookmarkControllerImpl) CreateBookmark(c *gin.Context) {
	userId := h.GetUser(c).UUID
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		bc.AbortHttp(c, bc.NewInvalidObjectIdError())
		return
	}

	var post post.Post
	if err := bc.PostRepo.FindById(context.Background(), postId, &post); err != nil {
		bc.AbortHttp(c, err)
		return
	}

	var bookmark bookmark.Bookmark
	if err := bc.Repo.FindOne(context.Background(), bson.M{"postId": postId, "userId": userId}, &bookmark); err != nil {
		if err != mongo.ErrNoDocuments || strings.ToLower(err.Error()) != "data not found" {
			bc.AbortHttp(c, err)
			return
		}
	} else {
		bc.AbortHttp(c, bc.New409Error("Conflict"))
		return
	}

	data := bc.Service.CreatePayload(postId, userId)
	if err := bc.Repo.CreateOne(context.Background(), &data); err != nil {
		bc.AbortHttp(c, err)
		return
	}

	bc.Write201Response(c, "success", data)
}

func (bc *BookmarkControllerImpl) DeleteBookmark(c *gin.Context) {
	bookmarkId, err := primitive.ObjectIDFromHex(c.Param("bookmarkId"))
	if err != nil {
		bc.AbortHttp(c, bc.NewInvalidObjectIdError())
		return
	}

	var data bookmark.Bookmark
	if err := bc.Repo.FIndById(context.Background(), bookmarkId, &data); err != nil {
		bc.AbortHttp(c, err)
		return
	}

	if err := bc.Repo.DeleteOneById(context.Background(), data.Id); err != nil {
		bc.AbortHttp(c, err)
		return
	}

	bc.Write200Response(c, "success", nil)
}
