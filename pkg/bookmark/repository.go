package bookmark

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewBookMarkRepo() BookmarkRepo {
	return &BookmarkRepoImpl{b.NewBaseRepo(b.GetCollection(b.Bookmark))}
}

func (r *BookmarkRepoImpl) CreateOne(ctx context.Context, data *Bookmark) error {
	if result, err := r.Create(ctx, data); err != nil {
		return err
	} else {
		data.Id = result
	}
	return nil
}

func (r *BookmarkRepoImpl) FindOne(ctx context.Context, query any, result *Bookmark) error {
	return r.FindOneByQuery(ctx, query, result)
}

func (r *BookmarkRepoImpl) FIndById(ctx context.Context, id primitive.ObjectID, result *Bookmark) error {
	return r.FindOneById(ctx, id, result)
}

func (r *BookmarkRepoImpl) DeleteOneById(ctx context.Context, id primitive.ObjectID) error {
	return r.BaseRepo.DeleteOneById(ctx, id)
}
