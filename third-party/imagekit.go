package thirdparty

import (
	"context"
	"os"

	"github.com/codedius/imagekit-go"
	"github.com/forumGamers/Octo-Cat/errors"
)

func NewImageKit() ImagekitService {
	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEKIT_PUBLIC_KEY"),
		PrivateKey: os.Getenv("IMAGEKIT_PRIVATE_KEY"),
	}

	ik, err := imagekit.NewClient(&opts)
	errors.PanicIfError(err)

	return &ImagekitServiceImpl{ik}
}

func (ik *ImagekitServiceImpl) UploadFile(ctx context.Context, upload UploadFile) (*imagekit.UploadResponse, error) {
	return ik.Client.Upload.ServerUpload(ctx, &imagekit.UploadRequest{
		File:              upload.Data,
		FileName:          upload.Name,
		UseUniqueFileName: true,
		Folder:            upload.Folder,
	})
}
