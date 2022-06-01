package aws

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	awshandlers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers"

	"github.com/kory-jp/react_golang_mux/api/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	Session *session.Session
}

func NewS3() *S3 {
	creds := credentials.NewStaticCredentials(config.Config.AwsAccessKeyID, config.Config.AwsSecretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(config.Config.AwsRegion),
	})
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	s3 := new(S3)
	s3.Session = sess
	return s3
}

type S3Result struct {
	Result *s3manager.UploadOutput
}

func (s3 *S3) ImageUploader(r *http.Request) (awshandlers.Result, error) {
	res := S3Result{}
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	if file, fileHeader, err = r.FormFile("image"); err != nil {
		if err == http.ErrMissingFile {
			fmt.Println("画像が投稿されていません")
			return nil, nil
		} else if err != nil {
			fmt.Println(err)
			log.Println(err)
			err = errors.New("画像の取り込み失敗しました")
			return nil, err
		}
	}
	defer file.Close()

	uploadFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileHeader.Filename)
	uploader := s3manager.NewUploader(s3.Session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Config.AwsBucket),
		Key:    aws.String(uploadFileName),
		Body:   file,
	})
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		err = errors.New("画像の保存に失敗しました")
		return nil, err
	}
	res.Result = result
	return res, nil
}

func (r S3Result) Location() string {
	return r.Result.Location
}

func (r S3Result) VersionID() *string {
	return r.Result.VersionID
}

func (r S3Result) UploadID() string {
	return r.Result.UploadID
}

func (r S3Result) ETag() *string {
	return r.Result.ETag
}
