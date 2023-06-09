package storage

import (
	"sthl/config"
	"sthl/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type S3Client struct {
	logger     *zap.Logger
	s3svc      *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func NewS3Client(logger *zap.Logger, cfg *config.Config) *S3Client {
	keyId, err := cfg.GetAwsAccessKeyId()
	if err != nil {
		return nil
	}
	secret, err := cfg.GetAwsSecretAccessKey()
	if err != nil {
		return nil
	}
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(keyId, secret, ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(cfg.GetAwsRegion()),
		Endpoint:         aws.String(cfg.GetS3Path()),
	})
	if err != nil {
		logger.Info("Failed to session.NewSession", zap.Error(err))
		return nil
	}

	s3Client := s3.New(sess)
	s3Uploader := s3manager.NewUploader(sess)
	s3Downloader := s3manager.NewDownloader(sess)

	// GetS3Path() == "" represent using default generated endpoint by aws sdk.
	// GetS3Path() != "" represent using localstack s3, create default backet every time
	if cfg.GetS3Path() != "" {
		resp, err := s3Client.CreateBucket(&s3.CreateBucketInput{
			ACL:    aws.String(s3.BucketCannedACLPublicRead),
			Bucket: aws.String(constants.S3BucketName),
		})
		if err != nil {
			logger.Info("Failed to CreateBucket", zap.Error(err))
		}
		logger.Info("localstack CreateBucket resp:", zap.Any("resp", resp))
	}

	listResult, err := s3Client.ListBuckets(nil)
	if err != nil {
		logger.Info("ListBuckets err:", zap.Error(err))
	}
	logger.Info("ListBuckets listResult:", zap.Any("listResult", listResult))

	return &S3Client{
		logger:     logger,
		s3svc:      s3Client,
		uploader:   s3Uploader,
		downloader: s3Downloader,
	}
}

func (s *S3Client) GetUploader() *s3manager.Uploader {
	return s.uploader
}
func (s *S3Client) GetDownloader() *s3manager.Downloader {
	return s.downloader
}
