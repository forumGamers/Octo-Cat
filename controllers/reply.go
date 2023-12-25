package controllers

import (
	"context"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/reply"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewReplyController(
	w web.ResponseWriter,
	r web.RequestReader,
	service reply.ReplyService,
	commentRepo comment.CommentRepo,
	validate *validator.Validate,
) ReplyController {
	return &ReplyControllerImpl{w, r, service, commentRepo, validate}
}

func (rc *ReplyControllerImpl) AddReply(c *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(c.Param("commentId"))
	if err != nil {
		rc.AbortHttp(c, rc.NewInvalidObjectIdError())
		return
	}

	var data web.CommentForm
	rc.GetParams(c, &data)

	if err := rc.Validator.Struct(&data); err != nil {
		rc.HttpValidationErr(c, err)
		return
	}

	var comment comment.Comment
	if err := rc.CommentRepo.FindById(context.Background(), commentId, &comment); err != nil {
		rc.AbortHttp(c, err)
		return
	}

	reply := rc.Service.CreatePayload(data, h.GetUser(c).UUID)
	if err := rc.CommentRepo.CreateReply(context.Background(), comment.Id, &reply); err != nil {
		rc.AbortHttp(c, err)
		return
	}

	rc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "success",
		Data:    reply,
	})
}

func (rc *ReplyControllerImpl) DeleteReply(c *gin.Context) {
	replyId, err := primitive.ObjectIDFromHex(c.Param("replyId"))
	if err != nil {
		rc.AbortHttp(c, rc.NewInvalidObjectIdError())
		return
	}
	commentId, err := primitive.ObjectIDFromHex(c.Param("commentId"))
	if err != nil {
		rc.AbortHttp(c, rc.NewInvalidObjectIdError())
		return
	}

	var reply comment.ReplyComment
	if err := rc.CommentRepo.FindReplyById(context.Background(), commentId, replyId, &reply); err != nil {
		rc.AbortHttp(c, err)
		return
	}

	if err := rc.Service.AuthorizeDeleteReply(reply, h.GetUser(c)); err != nil {
		rc.AbortHttp(c, err)
		return
	}

	if err := rc.CommentRepo.DeleteOneReply(context.Background(), commentId, replyId); err != nil {
		rc.AbortHttp(c, err)
		return
	}

	rc.WriteResponse(c, web.WebResponse{
		Code:    200,
		Message: "success",
	})
}
