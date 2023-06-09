package repository

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/ent/product"
	"sthl/storage"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IProductRepository interface {
	WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error
	CreateProduct(ctx context.Context, client *ent.Client, userId string, payload *dto.CreateProductDtoMappedDto) (*ent.Product, error)
	GetProductsTotalByUserId(ctx context.Context, client *ent.Client, userId string) (int, error)
	GetProducts(ctx context.Context, client *ent.Client, userId string, payload *dto.QueryProductsDto) (*dto.QueryProductsResponseDto, error)
	GetProductById(ctx context.Context, client *ent.Client, productId string) (*ent.Product, error)
	UpdateProductById(ctx context.Context, client *ent.Client, productId string, payload *dto.UpdateProductDto) (*ent.Product, error)
	SoftDeleteProductById(ctx context.Context, client *ent.Client, productId string) (*ent.Product, error)
}

type ProductRepository struct {
	logger *zap.Logger
}

func NewProductRepository(logger *zap.Logger) IProductRepository {
	return &ProductRepository{
		logger: logger,
	}
}
func (productRepo *ProductRepository) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTx(ctx, productRepo.logger, client, fn)
}

// CreateProduct
func (productRepo *ProductRepository) CreateProduct(
	ctx context.Context, client *ent.Client, userId string, payload *dto.CreateProductDtoMappedDto) (*ent.Product, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		productRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	result, err := client.Product.Create().
		SetUserID(userUuid).
		SetName(*payload.Name).
		SetPrice(*payload.Price).
		SetQuantity(*payload.Quantity).
		SetDescription(*payload.Description).
		SetStatus(*payload.Status).
		SetImgURL(*payload.ImgUrl).
		Save(ctx)
	if err != nil {
		productRepo.logger.Info("fail to client.Product.Create", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// GetProductsTotal
func (productRepo *ProductRepository) GetProductsTotalByUserId(ctx context.Context, client *ent.Client, userId string) (int, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		productRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return 0, constants.ErrBadRequest
	}

	total, err := client.Product.Query().Where(product.UserID(userUuid)).Count(ctx)
	if err != nil {
		productRepo.logger.Info("fail to count total", zap.Error(err))
		return 0, handleEntRepoErr(err)
	}
	return total, nil
}

// GetProducts
func (productRepo *ProductRepository) GetProducts(
	ctx context.Context, client *ent.Client, userId string, payload *dto.QueryProductsDto) (*dto.QueryProductsResponseDto, error) {

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		productRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit

	total, err := client.Product.Query().Where(product.UserID(userUuid)).Count(ctx)
	if err != nil {
		productRepo.logger.Info("fail to count total", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	// call ent client to Query
	result, err := client.Product.Query().
		Where(product.UserID(userUuid), product.NameContains(payload.Query)).
		Order(ent.Desc(product.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		productRepo.logger.Info("fail to client.Product.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	pagingResp := dto.NewPagingResponse(page, limit, total)
	data := dto.NewQueryProductsResponseDto(result, *pagingResp)
	return data, nil
}

// GetProductById
func (productRepo *ProductRepository) GetProductById(
	ctx context.Context, client *ent.Client, productId string) (*ent.Product, error) {

	productUuid, err := uuid.Parse(productId)
	if err != nil {
		productRepo.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Product.Get(ctx, productUuid)
	if err != nil {
		productRepo.logger.Info("fail to client.Product.Get", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// UpdateProductById
func (productRepo *ProductRepository) UpdateProductById(
	ctx context.Context, client *ent.Client, productId string, payload *dto.UpdateProductDto) (*ent.Product, error) {
	productUuid, err := uuid.Parse(productId)
	if err != nil {
		productRepo.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	result, err := client.Product.UpdateOneID(productUuid).
		SetName(*payload.Name).
		SetPrice(*payload.Price).
		SetQuantity(*payload.Quantity).
		SetDescription(*payload.Description).
		SetStatus(*payload.Status).
		SetImgURL(*payload.ImgUrl).
		Save(ctx)

	if err != nil {
		productRepo.logger.Info("fail to client.Product.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// SoftDeleteProductById
func (productRepo *ProductRepository) SoftDeleteProductById(
	ctx context.Context, client *ent.Client, productId string) (*ent.Product, error) {
	productUuid, err := uuid.Parse(productId)
	if err != nil {
		productRepo.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Product.UpdateOneID(productUuid).
		SetIsArchived(true).
		Save(ctx)

	if err != nil {
		productRepo.logger.Info("fail to client.Product.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}
