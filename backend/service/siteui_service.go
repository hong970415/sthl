package service

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ISiteUiService interface {
	// public
	GetSiteUiByUserId(ctx context.Context, userId string) (*ent.Siteui, error)
	// private
	UpsertSiteUiByUserId(ctx context.Context, userId string, payload *dto.UpsertSiteUiDto) (bool, error)
}
type SiteUiService struct {
	logger     *zap.Logger
	client     *ent.Client
	userRepo   repository.IUserRepository
	siteuiRepo repository.ISiteUiRepository
}

func NewSiteUiService(logger *zap.Logger, client *ent.Client,
	userRepo repository.IUserRepository, siteuiRepo repository.ISiteUiRepository) ISiteUiService {
	return &SiteUiService{
		logger:     logger,
		client:     client,
		userRepo:   userRepo,
		siteuiRepo: siteuiRepo,
	}
}

// GetSiteUiByUserId
func (siteuiSvc *SiteUiService) GetSiteUiByUserId(
	ctx context.Context, userId string) (*ent.Siteui, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		siteuiSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to GetProducts
	result, err := siteuiSvc.siteuiRepo.GetSiteUiByUserId(ctx, siteuiSvc.client, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpsertSiteUiByUserId
func (siteuiSvc *SiteUiService) UpsertSiteUiByUserId(
	ctx context.Context, userId string, payload *dto.UpsertSiteUiDto) (bool, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		siteuiSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}
	err = payload.Validate()
	if err != nil {
		siteuiSvc.logger.Info("fail to validate", zap.Error(err))
		return false, constants.ErrBadRequest
	}

	// upsert with transaction
	var result bool = false
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to UpsertSiteUiByUserId
		upsertResult, err := siteuiSvc.siteuiRepo.UpsertSiteUiByUserId(ctx, txc, userId, payload)
		result = upsertResult
		if err != nil {
			return err
		}
		return nil
	}

	err = siteuiSvc.siteuiRepo.WithTx(ctx, siteuiSvc.client, txFunc)
	if err != nil {
		return result, err
	}
	return result, nil
}
