package thirdparty

import (
	"context"

	"github.com/codedius/imagekit-go"
)

type ImagekitService interface {
	UploadFile(ctx context.Context, upload UploadFile) (*imagekit.UploadResponse, error)
}

type ImagekitServiceImpl struct {
	Client *imagekit.Client
}
