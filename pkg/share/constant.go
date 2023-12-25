package share

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShareRepo interface {
	DeleteMany(ctx context.Context, postId primitive.ObjectID) error
}

type ShareRepoImpl struct {
	b.BaseRepo
}
