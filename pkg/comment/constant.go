package comment

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"github.com/forumGamers/Octo-Cat/pkg/user"
	"github.com/forumGamers/Octo-Cat/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepo interface {
	CreateComment(ctx context.Context, data *Comment) error
	CreateReply(ctx context.Context, id primitive.ObjectID, data *ReplyComment) error
	FindById(ctx context.Context, id primitive.ObjectID, data *Comment) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	CreateMany(ctx context.Context, datas []any) (*mongo.InsertManyResult, error)
	DeleteReplyByPostId(ctx context.Context, postId primitive.ObjectID) error
	FindReplyById(ctx context.Context, id, replyId primitive.ObjectID, data *ReplyComment) error
	DeleteOneReply(ctx context.Context, id, replyId primitive.ObjectID) error
	DeleteMany(ctx context.Context, postId primitive.ObjectID) error
}

type CommentRepoImpl struct {
	b.BaseRepo
}

type CommentService interface {
	CreatePayload(data web.CommentForm, postId primitive.ObjectID, userId string) Comment
	AuthorizeDeleteComment(data Comment, user user.User) error
	InsertManyAndBindIds(ctx context.Context, datas []Comment) error
}

type CommentServiceImpl struct {
	Repo CommentRepo
}
