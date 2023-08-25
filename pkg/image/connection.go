package image

import (
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func ConnectS3(cfg config.Config) (*s3manager.Uploader,error) {

	var uploader *s3manager.Uploader 

	// create aws session
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(cfg.AWSS3REGION),
			Credentials: credentials.NewStaticCredentials(
				cfg.AWSS3ACCESSKEY,
				cfg.AWSS3SECRECTKEY,
				"",
			),
		},
	})

	if err != nil {
		return nil,err
	}

	// create an uploader
	uploader = s3manager.NewUploader(awsSession)

	return uploader,nil

}