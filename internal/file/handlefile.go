package file

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func HandleFile(formfile *multipart.Form, username string) *multipart.FileHeader {
	image := formfile.File["image"][0]
	imageExt := strings.Split(image.Filename, ".")
	image.Filename = fmt.Sprintf("%s.%s", username, imageExt)
	return image
}