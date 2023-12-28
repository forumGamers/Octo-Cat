package web

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewRequestReader() RequestReader {
	return &RequestReaderImpl{}
}

func (r *RequestReaderImpl) GetParams(c *gin.Context, p any) error {
	return c.ShouldBind(p)
}

type PostForm struct {
	File         []*multipart.FileHeader `form:"files[]" binding:"required"`
	Text         string                  `form:"text" json:"text" validate:"max=400"`
	AllowComment bool                    `form:"allowComment" json:"allowComment"`
	Privacy      string                  `form:"privacy" json:"privacy" validate:"oneof=Public FriendOnly Private"`
}

type CommentForm struct {
	Text string `json:"text" validate:"required" form:"text"`
}

type PostData struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId       string             `json:"userId" bson:"userId"`
	Text         string             `json:"text" bson:"text"`
	AllowComment bool               `json:"allowComment" bson:"allowComment"`
	Privacy      string             `json:"privacy" bson:"privacy"`
	Media        PostDataMedia      `json:"media" bson:"media"`
	CreatedAt    string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt    string             `json:"updatedAt" bson:"updatedAt"`
}

type PostDataMedia struct {
	Type string `json:"type" bson:"type"`
	Url  string `json:"url" bson:"url"`
	Id   string `json:"id" bson:"id"`
}

type PostDatas struct {
	Datas []PostData `json:"datas" binding:"required"`
}

type LikeData struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId    string             `json:"userId" bson:"userId"`
	PostId    primitive.ObjectID `json:"postId" bson:"postId"`
	CreatedAt string             `json:"CreatedAt" bson:"createdAt"`
	UpdatedAt string             `json:"UpdatedAt" bson:"updatedAt"`
}

type LikeDatas struct {
	Datas []LikeData `json:"datas" binding:"required"`
}

type CommentData struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Text      string             `json:"text" bson:"text"`
	UserId    string             `json:"userId" bson:"userId"`
	PostId    primitive.ObjectID `json:"postId" bson:"postId"`
	CreatedAt string             `json:"createdAt" bson:"CreatedAt"`
	UpdatedAt string             `json:"updatedAt" bson:"UpdatedAt"`
}

type CommentDatas struct {
	Datas []CommentData `json:"datas" binding:"required"`
}
