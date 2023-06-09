package repository

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/ent/siteui"
	"sthl/storage"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ISiteUiRepository interface {
	WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error
	GetSiteUiByUserId(ctx context.Context, client *ent.Client, userId string) (*ent.Siteui, error)
	UpsertSiteUiByUserId(ctx context.Context, client *ent.Client, userId string, payload *dto.UpsertSiteUiDto) (bool, error)
}

type SiteUiRepository struct {
	logger *zap.Logger
}

func NewSiteUiRepository(logger *zap.Logger) ISiteUiRepository {
	return &SiteUiRepository{
		logger: logger,
	}
}
func (siteuiRepo *SiteUiRepository) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTx(ctx, siteuiRepo.logger, client, fn)
}

// GetSiteUiByUserId
func (siteuiRepo *SiteUiRepository) GetSiteUiByUserId(
	ctx context.Context, client *ent.Client, userId string) (*ent.Siteui, error) {

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		siteuiRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Siteui.Query().
		Where(siteui.UserID(userUuid)).
		First(ctx)
	if err != nil {
		siteuiRepo.logger.Info("fail to client.Siteui.Query()", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

func (siteuiRepo *SiteUiRepository) UpsertSiteUiByUserId(
	ctx context.Context, client *ent.Client, userId string,
	payload *dto.UpsertSiteUiDto,
) (bool, error) {

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		siteuiRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}

	err = client.Siteui.Create().
		SetUserID(userUuid).
		SetSitename(*payload.Sitename).
		SetHomepageImgUrl(*payload.HomepageImgUrl).
		SetHomepageText(*payload.HomepageText).
		SetHomepageTextColor(*payload.HomepageTextColor).
		OnConflict(
			sql.ConflictColumns(siteui.FieldUserID),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		siteuiRepo.logger.Info("fail to client.Siteui.Query()", zap.Error(err))
		return false, handleEntRepoErr(err)
	}
	return true, nil
}
