package controllers

import (
	"context"
	"sync"
	"time"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/like"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewLikeController(
	w web.ResponseWriter,
	r web.RequestReader,
	service like.LikeService,
	repo like.LikeRepo,
	postRepo post.PostRepo,
	validate *validator.Validate,
) LikeController {
	return &LikeControllerImpl{w, r, service, repo, postRepo, validate}
}

func (lc *LikeControllerImpl) LikePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		lc.AbortHttp(c, lc.NewInvalidObjectIdError())
		return
	}

	id := h.GetUser(c).UUID
	var post post.Post
	if err := lc.PostRepo.FindOneById(context.Background(), postId, &post); err != nil {
		lc.AbortHttp(c, err)
		return
	}

	var data like.Like
	if err := lc.Repo.GetLikesByUserIdAndPostId(context.Background(), postId, id, &data); err != nil {
		//perbaiki nanti
		if err != nil {
			lc.AbortHttp(c, err)
			return
		}
	}

	if data.Id != primitive.NilObjectID {
		lc.AbortHttp(c, lc.New409Error("Conflict"))
		return
	}

	newLike := like.Like{
		UserId:    id,
		PostId:    post.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := lc.Repo.AddLikes(context.Background(), newLike)
	if err != nil {
		lc.AbortHttp(c, err)
		return
	}

	newLike.Id = result

	lc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "success",
		Data:    newLike,
	})
}

func (lc *LikeControllerImpl) UnlikePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		lc.AbortHttp(c, lc.NewInvalidObjectIdError())
		return
	}

	userId := h.GetUser(c).UUID
	var like like.Like

	if err := lc.Repo.GetLikesByUserIdAndPostId(context.Background(), postId, userId, &like); err != nil {
		lc.AbortHttp(c, err)
		return
	}

	if err := lc.Repo.DeleteLike(context.Background(), postId, userId); err != nil {
		lc.AbortHttp(c, err)
		return
	}

	lc.WriteResponse(c, web.WebResponse{
		Code:    200,
		Message: "success",
	})
}

func (lc *LikeControllerImpl) BulkLikes(c *gin.Context) {
	if h.GetStage(c) != "Development" {
		lc.CustomMsgAbortHttp(c, "No Content", 204)
		return
	}

	var datas web.LikeDatas
	lc.GetParams(c, &datas)

	var likes []like.Like
	var wg sync.WaitGroup
	for _, data := range datas.Datas {
		wg.Add(1)
		go func(data web.LikeData) {
			defer wg.Done()
			postId, _ := primitive.ObjectIDFromHex(data.PostId.Hex())
			t, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.CreatedAt)
			u, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.UpdatedAt)
			likes = append(likes, like.Like{
				PostId:    postId,
				UserId:    data.UserId,
				CreatedAt: t,
				UpdatedAt: u,
			})
		}(data)
	}

	wg.Wait()
	lc.Service.InsertManyAndBindIds(context.Background(), likes)

	lc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "Success",
		Data:    likes,
	})
}
