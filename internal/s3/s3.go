package s3

import (
	"habitat/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateSession(awsConfig config.AWSConfig) *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKeyId, awsConfig.AccessKeySecret, ""),
	}))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}
