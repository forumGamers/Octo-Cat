package bookmark

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookmarkRepo interface {
	CreateOne(ctx context.Context, data *Bookmark) error
	FindOne(ctx context.Context, query any, result *Bookmark) error
	FIndById(ctx context.Context, id primitive.ObjectID, result *Bookmark) error
	DeleteOneById(ctx context.Context, id primitive.ObjectID) error
}

type BookmarkRepoImpl struct {
	b.BaseRepo
}

type BookmarkService interface {
	CreatePayload(postId primitive.ObjectID, userId string) Bookmark
}

type BookmarkServiceImpl struct {
	Repo BookmarkRepo
}
