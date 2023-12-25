package reply

import (
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/user"
	"github.com/forumGamers/Octo-Cat/web"
)

type ReplyService interface {
	CreatePayload(data web.CommentForm, userId string) comment.ReplyComment
	AuthorizeDeleteReply(data comment.ReplyComment, user user.User) error
}

type ReplyServiceImpl struct {
	Repo comment.CommentRepo
}
