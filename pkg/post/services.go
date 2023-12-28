package post

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/forumGamers/Octo-Cat/errors"
	h "github.com/forumGamers/Octo-Cat/helpers"
	tp "github.com/forumGamers/Octo-Cat/third-party"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewPostService(repo PostRepo, ik tp.ImagekitService) PostService {
	return &PostServiceImpl{repo, ik}
}

func (s *PostServiceImpl) InsertManyAndBindIds(ctx context.Context, datas []Post) error {
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

func (s *PostServiceImpl) UploadPostMedia(postMedias *[]Media, file *multipart.FileHeader, c *gin.Context) error {
	media, savedFile, err := h.SaveUploadedFile(c, file)
	if err != nil {
		return errors.NewError("Failed to process file", 501)
	}
	defer os.Remove(h.GetUploadDir(file.Filename))
	defer savedFile.Close()

	fileType, err := h.CheckFileType(file)
	if err != nil {
		return errors.NewError(err.Error(), 400)
	}

	response, err := s.Ik.UploadFile(context.Background(), tp.UploadFile{
		Data:   media,
		Name:   savedFile.Name(),
		Folder: fmt.Sprintf("post_%s", fileType),
	})
	if err != nil {
		return errors.NewError("Failed to save file", 502)
	}

	*postMedias = append(*postMedias, Media{
		Url:  response.URL,
		Id:   response.FileID,
		Type: fileType,
	})

	return nil
}

func (s *PostServiceImpl) GetPostTags(text string) []string {
	modified := text
	for _, p := range []rune("!@#$%^&*)(_=+?.,;:'") {
		modified = strings.ReplaceAll(modified, string(p), " ")
	}
	return strings.Split(modified, " ")
}

func (s *PostServiceImpl) CreatePostPayload(userId, text, privacy string, allowComment bool, media []Media, tags []string) Post {
	return Post{
		UserId:       userId,
		Text:         text,
		Media:        media,
		AllowComment: allowComment,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Tags:         tags,
		Privacy:      privacy,
	}
}
