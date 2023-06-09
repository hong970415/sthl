package repository

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/ent/imageinfo"
	"sthl/storage"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IImgInfoRepository interface {
	WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error
	CreateImg(ctx context.Context, client *ent.Client, userId string, payload *dto.CreateImgDto) (*ent.Imageinfo, error)
	CheckImgNameExist(ctx context.Context, client *ent.Client, userId string, imgName string) (bool, error)
	GetImgByUserId(ctx context.Context, client *ent.Client, userId string, imgInfoId int) (*ent.Imageinfo, error)
	GetImgsByUserId(ctx context.Context, client *ent.Client, userId string, payload *dto.QueryImgsInfoDto) (*dto.QueryImgsInfoResponseDto, error)
	UpdateImgInfoById(ctx context.Context, client *ent.Client, imgInfoId int, payload *dto.UpdateImgInfoDto) (*ent.Imageinfo, error)
}

type ImgInfoRepository struct {
	logger *zap.Logger
}

func NewImgInfoRepository(logger *zap.Logger) IImgInfoRepository {
	return &ImgInfoRepository{
		logger: logger,
	}
}
func (imginfoRepo *ImgInfoRepository) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTx(ctx, imginfoRepo.logger, client, fn)
}

// CreateImg
func (imginfoRepo *ImgInfoRepository) CreateImg(ctx context.Context, client *ent.Client,
	userId string, payload *dto.CreateImgDto) (*ent.Imageinfo, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		imginfoRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Imageinfo.Create().
		SetUserID(userUuid).
		SetImgS3IDKey(*payload.ImgS3IdKey).
		SetImgName(*payload.ImgName).
		SetImgURL(*payload.ImgURL).
		SetImgSize(*payload.ImgSize).
		Save(ctx)
	if err != nil {
		imginfoRepo.logger.Info("fail to client.Imageinfo.Create", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}
func (imginfoRepo *ImgInfoRepository) CheckImgNameExist(ctx context.Context, client *ent.Client,
	userId string, imgName string) (bool, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		imginfoRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}

	result, err := client.Imageinfo.Query().
		Where(imageinfo.UserID(userUuid), imageinfo.ImgNameEQ(imgName)).
		Exist(ctx)
	if err != nil {
		imginfoRepo.logger.Info("fail to client.Imageinfo.Query", zap.Error(err))
		return false, handleEntRepoErr(err)
	}
	return result, nil
}

// GetImgByUserId
func (imginfoRepo *ImgInfoRepository) GetImgByUserId(
	ctx context.Context, client *ent.Client, userId string, imgInfoId int) (*ent.Imageinfo, error) {
	_, err := uuid.Parse(userId)
	if err != nil {
		imginfoRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Imageinfo.Get(ctx, imgInfoId)
	if err != nil {
		imginfoRepo.logger.Info("fail to client.Imageinfo.Get", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

func (imginfoRepo *ImgInfoRepository) GetImgsTotalByUserId(ctx context.Context, client *ent.Client, userId string) (int, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		imginfoRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return 0, constants.ErrBadRequest
	}

	total, err := client.Imageinfo.Query().Where(imageinfo.UserID(userUuid)).Count(ctx)
	if err != nil {
		imginfoRepo.logger.Info("fail to count total", zap.Error(err))
		return 0, handleEntRepoErr(err)
	}
	return total, nil
}

// GetImgsByUserId
func (imginfoRepo *ImgInfoRepository) GetImgsByUserId(
	ctx context.Context, client *ent.Client, userId string, payload *dto.QueryImgsInfoDto) (*dto.QueryImgsInfoResponseDto, error) {

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		imginfoRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit

	total, err := client.Imageinfo.Query().Where(imageinfo.UserID(userUuid)).Count(ctx)
	if err != nil {
		imginfoRepo.logger.Info("fail to count total", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	// call ent client to Query
	result, err := client.Imageinfo.Query().
		Where(imageinfo.UserID(userUuid), imageinfo.ImgNameContains(payload.Query)).
		Order(ent.Desc(imageinfo.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		imginfoRepo.logger.Info("fail to client.Imageinfo.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	pagingResp := dto.NewPagingResponse(page, limit, total)
	data := dto.NewQueryImgsInfoResponseDto(result, *pagingResp)
	return data, nil
}

// UpdateImgInfoById
func (imginfoRepo *ImgInfoRepository) UpdateImgInfoById(
	ctx context.Context, client *ent.Client, imgInfoId int, payload *dto.UpdateImgInfoDto) (*ent.Imageinfo, error) {
	result, err := client.Imageinfo.UpdateOneID(imgInfoId).
		SetImgSize(*payload.ImgSize).
		Save(ctx)

	if err != nil {
		imginfoRepo.logger.Info("fail to client.Imageinfo.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}
