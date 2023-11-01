package s3

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadObject(bucket string, filePath string, fileName string, sess *session.Session) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable To Open File")
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err == nil {
		fmt.Printf("Successfuully upload %q to %q\n", fileName, bucket)
	} else {
		return errors.New("Unable to uploead file")
	}

	return nil
}
