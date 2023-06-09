package service

import (
	"context"
	"errors"
	"net/http"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/repository"
	"sthl/storage"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IAlbumService interface {
	// private
	UploadFile(r *http.Request) (*ent.Imageinfo, error)
	GetImgsByUserId(ctx context.Context, userId string, payload *dto.QueryImgsInfoDto) (*dto.QueryImgsInfoResponseDto, error)
	UpdateS3ImageDataById(r *http.Request) (*ent.Imageinfo, error)
}
type AlbumService struct {
	logger      *zap.Logger
	entClient   *ent.Client
	s3Client    *storage.S3Client
	imginfoRepo repository.IImgInfoRepository
}

func NewAlbumService(logger *zap.Logger, entClient *ent.Client,
	s3Client *storage.S3Client, imginfoRepo repository.IImgInfoRepository) IAlbumService {
	return &AlbumService{
		logger:      logger,
		entClient:   entClient,
		s3Client:    s3Client,
		imginfoRepo: imginfoRepo,
	}
}

// UploadFile
func (gallerySvc *AlbumService) UploadFile(r *http.Request) (*ent.Imageinfo, error) {
	// extract ctx from request
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		gallerySvc.logger.Info("fail to extract userInfo from ctx")
		return nil, constants.ErrInternalServer
	}

	err := r.ParseMultipartForm(constants.MaxFileSize)
	if err != nil {
		gallerySvc.logger.Info("ParseMultipartForm exceeded", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	key := "file"
	file, header, err := r.FormFile(key)
	if errors.Is(err, http.ErrMissingFile) {
		return nil, constants.ErrBadRequest
	}

	gallerySvc.logger.Info(key,
		zap.String("name", header.Filename), zap.Int64("size", header.Size), zap.Bool("larger than 4mb", header.Size > constants.MaxFileSize))
	if header.Size > constants.MaxFileSize {
		return nil, constants.ErrBadRequest
	}
	nameExist, err := gallerySvc.imginfoRepo.CheckImgNameExist(ctx, gallerySvc.entClient, authenticatedUserInfo, header.Filename)
	if err != nil {
		return nil, err
	}
	if nameExist {
		gallerySvc.logger.Info("filename already exist")
		return nil, constants.ErrBadRequest
	}

	// s3 upload
	idKey := authenticatedUserInfo + "/" + uuid.NewString()
	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(constants.S3BucketName),
		Key:         aws.String(idKey),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: aws.String(header.Header.Get("content-type")),
		Body:        file,
	}
	s3result, err := gallerySvc.s3Client.GetUploader().Upload(uploadInput)
	if err != nil {
		return nil, err
	}

	// create with transaction
	var result *ent.Imageinfo
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to CreateImg
		data := dto.NewCreateImgDto(&header.Filename, &s3result.Location, &header.Size, &idKey)
		createResult, err := gallerySvc.imginfoRepo.CreateImg(ctx, txc, authenticatedUserInfo, data)
		result = createResult
		if err != nil {
			return err
		}
		return nil
	}

	err = gallerySvc.imginfoRepo.WithTx(ctx, gallerySvc.entClient, txFunc)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetImgsByUserId
func (gallerySvc *AlbumService) GetImgsByUserId(ctx context.Context, userId string, payload *dto.QueryImgsInfoDto) (*dto.QueryImgsInfoResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		gallerySvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	ensuredPaging := payload.Ensure()
	ensuredPayload := dto.NewQueryImgsInfoDto(*ensuredPaging)

	// call repo to GetImgsByUserId
	result, err := gallerySvc.imginfoRepo.GetImgsByUserId(ctx, gallerySvc.entClient, userId, ensuredPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func (gallerySvc *AlbumService) UpdateS3ImageDataById(ctx context.Context, userId string, payload *dto.UpdateImgInfoDto) (*ent.Imageinfo, error) {
func (gallerySvc *AlbumService) UpdateS3ImageDataById(r *http.Request) (*ent.Imageinfo, error) {
	// extract ctx from request
	ctx := r.Context()

	imgInfoIdParam, err := strconv.Atoi(chi.URLParam(r, "imgInfoId"))
	if err != nil {
		gallerySvc.logger.Info("fail to  strconv.Atoi")
		return nil, constants.ErrBadRequest
	}

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		gallerySvc.logger.Info("fail to extract userInfo from ctx")
		return nil, constants.ErrInternalServer
	}

	err = r.ParseMultipartForm(constants.MaxFileSize)
	if err != nil {
		gallerySvc.logger.Info("ParseMultipartForm exceeded", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// validate file size
	key := "file"
	file, header, err := r.FormFile(key)
	if errors.Is(err, http.ErrMissingFile) {
		return nil, constants.ErrBadRequest
	}
	gallerySvc.logger.Info(key,
		zap.String("name", header.Filename), zap.Int64("size", header.Size), zap.Bool("larger than 4mb", header.Size > constants.MaxFileSize))
	if header.Size > constants.MaxFileSize {
		return nil, constants.ErrBadRequest
	}

	// call repo to get imginfo
	imgInfoData, err := gallerySvc.imginfoRepo.GetImgByUserId(ctx, gallerySvc.entClient, authenticatedUserInfo, imgInfoIdParam)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	// validate img name exist
	nameExist := imgInfoData.ImgName == header.Filename
	if nameExist {
		gallerySvc.logger.Info("filename already exist")
		return nil, constants.ErrBadRequest
	}

	// s3 upload
	idKey := imgInfoData.ImgS3IDKey
	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(constants.S3BucketName),
		Key:         aws.String(idKey),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: aws.String(header.Header.Get("content-type")),
		Body:        file,
	}
	_, err = gallerySvc.s3Client.GetUploader().Upload(uploadInput)
	if err != nil {
		return nil, err
	}

	// update with transaction
	var result *ent.Imageinfo
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to UpdateImgInfoById
		data := dto.NewUpdateImgInfoDto(&header.Size)
		updateResult, err := gallerySvc.imginfoRepo.UpdateImgInfoById(ctx, txc, imgInfoIdParam, data)
		result = updateResult
		if err != nil {
			return err
		}
		return nil
	}

	err = gallerySvc.imginfoRepo.WithTx(ctx, gallerySvc.entClient, txFunc)
	if err != nil {
		return result, err
	}
	return result, nil
}
