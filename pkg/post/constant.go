package post

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepo interface {
	Create(ctx context.Context, data Post) error
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
	InsertManyAndBindIds(ctx context.Context, datas []Post) error
}

type PostServiceImpl struct {
	Repo PostRepo
}
