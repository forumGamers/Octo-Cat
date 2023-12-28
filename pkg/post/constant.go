package post

import (
	"context"
	"mime/multipart"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	tp "github.com/forumGamers/Octo-Cat/third-party"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepo interface {
	Create(ctx context.Context, data *Post) error
	FindById(ctx context.Context, id primitive.ObjectID, data *Post) error
	GetSession() (mongo.Session, error)
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	CreateMany(ctx context.Context, datas []any) (*mongo.InsertManyResult, error)
	FindOneById(ctx context.Context, id primitive.ObjectID, data any) error
}

type PostRepoImpl struct {
	b.BaseRepo
}

type PostService interface {
	UploadPostMedia(postMedias *[]Media, file *multipart.FileHeader, c *gin.Context) error
	InsertManyAndBindIds(ctx context.Context, datas []Post) error
	GetPostTags(text string) []string
	CreatePostPayload(userId, text, privacy string, allowComment bool, media []Media, tags []string) Post
}

type PostServiceImpl struct {
	Repo PostRepo
	Ik   tp.ImagekitService
}
