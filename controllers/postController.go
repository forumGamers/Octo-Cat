package controllers

import (
	"context"
	"fmt"
	"sync"
	"time"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/comment"
	"github.com/forumGamers/Octo-Cat/pkg/like"
	"github.com/forumGamers/Octo-Cat/pkg/post"
	"github.com/forumGamers/Octo-Cat/pkg/share"
	tp "github.com/forumGamers/Octo-Cat/third-party"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPostControllers(
	w web.ResponseWriter,
	r web.RequestReader,
	service post.PostService,
	repo post.PostRepo,
	commentRepo comment.CommentRepo,
	likeRepo like.LikeRepo,
	shareRepo share.ShareRepo,
	Ik tp.ImagekitService,
	validate *validator.Validate,
) PostController {
	return &PostControllerImpl{w, r, service, repo, commentRepo, likeRepo, shareRepo, Ik, validate}
}

func (pc *PostControllerImpl) CreatePost(c *gin.Context) {
	var form web.PostForm
	pc.GetParams(c, &form)

	if err := pc.Validator.Struct(&form); err != nil {
		pc.HttpValidationErr(c, err)
		return
	}

	formInput, err := c.MultipartForm()
	if err != nil {
		pc.AbortHttp(c, pc.New501Error("Cannot process form data"))
		return
	}

	var postMedias []post.Media
	files := formInput.File["file"]
	if len(files) > 0 {
		errCh := make(chan error)
		var wg sync.WaitGroup
		for _, file := range files {
			wg.Add(1)
			go func(postMedias *[]post.Media) {
				defer wg.Done()
				media, savedFile, err := h.SaveUploadedFile(c, file)
				if err != nil {
					errCh <- err
					return
				}

				foldername, err := h.CheckFileType(file)
				if err != nil {
					errCh <- err
					return
				}

				_, err = pc.Ik.UploadFile(context.Background(), tp.UploadFile{
					Data:   media,
					Name:   savedFile.Name(),
					Folder: fmt.Sprintf("post_%s", foldername),
				})
				if err != nil {
					errCh <- err
					return
				}

				errCh <- nil
			}(&postMedias)
		}

		wg.Wait()
		close(errCh)

		for err := range errCh {
			if err != nil {
				pc.AbortHttp(c, err)
				return
			}
		}
	}

	// user := h.GetUser(c)

	// if form.File != nil {
	// 	fileInfo.Media, fileInfo.SavedFile, err := h.SaveUploadedFile(c, form.File)
	// 	if err != nil {
	// 		pc.New400Error(err.Error())
	// 	}

	// 	fileInfo.FolderName, err = h.CheckFileType(form.File)
	// 	if err != nil {
	// 		pc.New400Error(err.Error())
	// 	}

	// 	fileInfo.FolderName = "post_" + fileInfo.FolderName
	// 	fileInfo.FileName = form.File.Filename
	// }

	// data, err := pc.Service.CreatePostPayload(context.Background(), &form, user, tp.UploadFile{
	// 	Data:   fileInfo.Media,
	// 	Folder: fileInfo.FolderName,
	// 	Name:   fileInfo.FileName,
	// })

	// if err != nil {
	// 	pc.AbortHttp(c, err)
	// 	return
	// }

	// pc.Repo.Create(context.Background(), &data)

	// if form.File != nil {
	// 	fileInfo.SavedFile.Close()
	// 	os.Remove(h.GetUploadDir(fileInfo.FileName))
	// }

	// data.Text = h.Decryption(data.Text)

	// pc.WriteResponse(c, web.WebResponse{
	// 	Code:    201,
	// 	Message: "Success",
	// 	Data:    data,
	// })
}

func (pc *PostControllerImpl) DeletePost(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		pc.AbortHttp(c, pc.NewInvalidObjectIdError())
		return
	}

	var data post.Post
	if err := pc.Repo.FindById(context.Background(), postId, &data); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	user := h.GetUser(c)

	if data.UserId != user.UUID || user.LoggedAs != "Admin" {
		pc.AbortHttp(c, pc.New403Error("Forbidden"))
		return
	}

	session, err := pc.Repo.GetSession()
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	defer session.EndSession(context.Background())

	if err := session.StartTransaction(); err != nil {
		pc.AbortHttp(c, err)
		return
	}

	ctx := mongo.NewSessionContext(context.Background(), session)
	var wg sync.WaitGroup
	errCh := make(chan error)
	wg.Add(4)
	runRountine := func(f func()) {
		defer wg.Done()
		f()
	}

	// go runRountine(func() {
	// 	pc.Service.DeletePostMedia(ctx, data, errCh)
	// })
	go runRountine(func() {
		errCh <- pc.LikeRepo.DeletePostLikes(ctx, data.Id)
	})
	go runRountine(func() {
		errCh <- pc.CommentRepo.DeleteMany(ctx, data.Id)
	})
	go runRountine(func() {
		errCh <- pc.Repo.DeleteOne(ctx, data.Id)
	})
	go runRountine(func() {
		errCh <- pc.ShareRepo.DeleteMany(ctx, data.Id)
	})

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			session.AbortTransaction(ctx)
			pc.AbortHttp(c, err)
			return
		}
	}

	if err := session.CommitTransaction(ctx); err != nil {
		session.AbortTransaction(ctx)
		pc.AbortHttp(c, err)
		return
	}

	pc.WriteResponse(c, web.WebResponse{
		Message: "success",
		Code:    200,
	})
}

func (pc *PostControllerImpl) BulkCreatePost(c *gin.Context) {
	if h.GetStage(c) != "Development" {
		pc.CustomMsgAbortHttp(c, "No Content", 204)
		return
	}

	var datas web.PostDatas
	c.ShouldBind(&datas)

	var posts []post.Post
	var wg sync.WaitGroup
	for _, data := range datas.Datas {
		wg.Add(1)
		go func(data web.PostData) {
			defer wg.Done()
			t, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.CreatedAt)
			u, _ := time.Parse("2006-01-02T15:04:05Z07:00", data.UpdatedAt)
			data.Text = h.Encryption(data.Text)
			posts = append(posts, post.Post{
				UserId: data.UserId,
				Text:   data.Text,
				Media: []post.Media{
					{
						Url:  data.Media.Url,
						Id:   data.Media.Id,
						Type: data.Media.Type,
					},
				},
				AllowComment: data.AllowComment,
				Tags:         []string{},
				Privacy:      data.Privacy,
				CreatedAt:    t,
				UpdatedAt:    u,
			})
		}(data)
	}
	wg.Wait()

	pc.Service.InsertManyAndBindIds(context.Background(), posts)

	pc.WriteResponse(c, web.WebResponse{
		Code:    201,
		Message: "Success",
		Data:    posts,
	})
}
