package controllers

import (
	"context"
	"sync"
	"time"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCommentController(
	w web.ResponseWriter,
	r web.RequestReader,
	repo comment.CommentRepo,
	service comment.CommentService,
	postRepo post.PostRepo,
	validate *validator.Validate,
) CommentController {
	return &CommentControllerImpl{w, r, repo, service, postRepo, validate}
}

func (pc *CommentControllerImpl) CreateComment(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		pc.AbortHttp(c, pc.NewInvalidObjectIdError())
		return
	}

	var data web.CommentForm
	pc.GetParams(c, &data)

	if err := pc.Validator.Struct(&data); err != nil {
		pc.HttpValidationErr(c, err)
		return
	}

	var post post.Post
	if err := pc.PostRepo.FindOneById(context.Background(), postId, &post); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	comment := pc.Service.CreatePayload(data, postId, h.GetUser(c).UUID)
	if err := pc.Repo.CreateComment(context.Background(), &comment); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	comment.Text = h.Decryption(comment.Text)

	pc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "success",
		Data:    comment,
	})
}

func (pc *CommentControllerImpl) DeleteComment(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("commentId"))
	if err != nil {
		pc.AbortHttp(c, pc.NewInvalidObjectIdError())
		return
	}

	var comment comment.Comment
	if err := pc.Repo.FindById(context.Background(), commentId, &comment); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	if err := pc.Service.AuthorizeDeleteComment(comment, h.GetUser(c)); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	if err := pc.Repo.DeleteOne(context.Background(), commentId); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.WriteResponse(c, web.WebResponse{
		Code:    200,
		Message: "success",
	})
}

func (pc *CommentControllerImpl) BulkComment(c *gin.Context) {
	if h.GetStage(c) != "Development" {
		pc.CustomMsgAbortHttp(c, "No Content", 204)
		return
	}

	var datas web.CommentDatas
	pc.GetParams(c, &datas)

	var comments []comment.Comment
	var wg sync.WaitGroup
	for _, data := range datas.Datas {
		wg.Add(1)
		go func(data web.CommentData) {
			defer wg.Done()
			postId, _ := primitive.ObjectIDFromHex(data.PostId.Hex())
			t, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.CreatedAt)
			u, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.UpdatedAt)
			comments = append(comments, comment.Comment{
				PostId:    postId,
				UserId:    data.UserId,
				CreatedAt: t,
				UpdatedAt: u,
				Text:      data.Text,
				Reply:     []comment.ReplyComment{},
			})
		}(data)
	}

	wg.Wait()
	pc.Service.InsertManyAndBindIds(context.Background(), comments)

	pc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "success",
		Data:    comments,
	})
}
