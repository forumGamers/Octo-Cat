package comment

import (
	"context"
	"time"

	"github.com/forumGamers/Octo-Cat/errors"
	"github.com/forumGamers/Octo-Cat/pkg/user"
	"github.com/forumGamers/Octo-Cat/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCommentService(repo CommentRepo) CommentService {
	return &CommentServiceImpl{repo}
}

func (s *CommentServiceImpl) CreatePayload(data web.CommentForm, postId primitive.ObjectID, userId string) Comment {
	return Comment{
		UserId:    userId,
		Text:      data.Text,
		PostId:    postId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *CommentServiceImpl) AuthorizeDeleteComment(data Comment, user user.User) error {
	//nanti yang punya post juga bisa hapus
	if user.UUID != data.UserId || user.LoggedAs != "Admin" {
		return errors.NewError("unauthorized", 401)
	}
	return nil
}

func (s *CommentServiceImpl) InsertManyAndBindIds(ctx context.Context, datas []Comment) error {
	var payload []any

	for _, data := range datas {
		payload = append(payload, data)
	}

	ids, err := s.Repo.CreateMany(ctx, payload)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids.InsertedIDs); i++ {
		id := ids.InsertedIDs[i].(primitive.ObjectID)
		datas[i].Id = id
	}
	return nil
}
