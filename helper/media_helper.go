package helper

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	config "github.com/thitiphongD/cloudinary-gin-api/configs"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(
		config.ENVCLOUNDNAME(),
		config.ENVCLOUNDAPIKEY(),
		config.ENVCLOUNDAPISECRET(),
	)

	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.ENVCLOUNDUPLOADFOLDER()})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
