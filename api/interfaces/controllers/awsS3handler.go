package awshandlers

import "net/http"

type S3 interface {
	ImageUploader(*http.Request) (Result, error)
}

type Result interface {
	Location() string
	VersionID() *string
	UploadID() string
	ETag() *string
}
