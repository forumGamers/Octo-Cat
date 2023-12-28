package thirdparty

import (
	"context"

	"github.com/codedius/imagekit-go"
)

type ImagekitService interface {
	UploadFile(ctx context.Context, upload UploadFile) (*imagekit.UploadResponse, error)
	DeleteFile(ctx context.Context, fileId string) error
	DeleteBulkFile(ctx context.Context, fileIds []string) error
}

type ImagekitServiceImpl struct {
	Client *imagekit.Client
}
