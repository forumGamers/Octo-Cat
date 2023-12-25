package reply

import (
	"time"

	"github.com/forumGamers/Octo-Cat/errors"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/user"
	"github.com/forumGamers/Octo-Cat/web"
)

func NewReplyService(repo comment.CommentRepo) ReplyService {
	return &ReplyServiceImpl{repo}
}

func (rs *ReplyServiceImpl) CreatePayload(data web.CommentForm, userId string) comment.ReplyComment {
	return comment.ReplyComment{
		UserId:    userId,
		Text:      data.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (rs *ReplyServiceImpl) AuthorizeDeleteReply(data comment.ReplyComment, user user.User) error {
	//nanti yang punya post juga bisa hapus
	if user.UUID != data.UserId || user.LoggedAs != "Admin" {
		return errors.NewError("unauthorized", 401)
	}
	return nil
}
