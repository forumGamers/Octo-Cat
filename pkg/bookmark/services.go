package bookmark

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewBookMarkService(r BookmarkRepo) BookmarkService {
	return &BookmarkServiceImpl{r}
}

func (s *BookmarkServiceImpl) CreatePayload(postId primitive.ObjectID, userId string) Bookmark {
	bookmark := Bookmark{
		PostId:    postId,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return bookmark
}
