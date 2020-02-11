package infra

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"AY1st/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// GetS3FileBucketName S3 ファイル バケット名
func GetS3FileBucketName() string {
	return fmt.Sprintf("gc-sun-%s-files", os.Getenv("ENVCODE"))
}

// GetS3LogBucketName S3 ログ バケット名
func GetS3LogBucketName() string {
	return fmt.Sprintf("gc-sun-%s-logs", os.Getenv("ENVCODE"))
}

// GetS3NsDataBucketName S3 NSデータ バケット名
func GetS3NsDataBucketName() string {
	return fmt.Sprintf("gc-sun-%s-data", os.Getenv("ENVCODE"))
}

// S3Client は S3 署名付きURL 発行
type S3Client struct {
	Client     *s3.S3
	Manager    *s3manager.Uploader
	BucketName string
}

// NewS3Client は S3 URL Signer
func NewS3Client(bucketName string) *S3Client {
	s := GetAWSSession()
	s3Client := &S3Client{
		Client:     s3.New(s),
		Manager:    s3manager.NewUploader(s),
		BucketName: bucketName,
	}

	return s3Client
}

// GetSignedURL は S3 Pre-Signed URL
func (c *S3Client) GetSignedURL(filepath string, expiresSec int) (string, error) {
	req, _ := c.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(filepath),
	})

	preSignedURL, err := req.Presign(time.Duration(expiresSec) * time.Second)
	if err != nil {
		util.GetLogger().Error(err)
		return "", err
	}

	return preSignedURL, nil
}

// UploadFile uploads file to specific path from io.Reader
func (c *S3Client) UploadFile(filepath string, b io.ReadSeeker) error {

	_, err := c.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(filepath),
		Body:   b,
	})

	return err
}

// UploadCSVToS3 is UploadCSVToS3
func UploadCSVToS3(b io.ReadSeeker, bucketName string) (string, error) {
	key := util.GetFormatedTimeNow() + ".csv"

	sess := NewS3Client(bucketName)

	_, err := sess.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(sess.BucketName),
		Key:    aws.String(key),
		Body:   b,
	})
	if err != nil {
		return "", err
	}

	// PreSignedURLを発行
	downloadPath, nil := sess.GetSignedURL(key, 300)
	if err != nil {
		return "", err
	}

	return downloadPath, nil
}

// UploadFile is
func UploadFile(file multipart.File, bucketName, contentType, key string) (path string, err error) {

	sess := NewS3Client(bucketName)

	uploaded, err := sess.Manager.Upload(&s3manager.UploadInput{
		Body:        file,
		Bucket:      aws.String(sess.BucketName),
		ContentType: aws.String(contentType),
		Key:         aws.String(key),
	})

	if err != nil {
		return "", err
	}

	return uploaded.Location, nil
}
