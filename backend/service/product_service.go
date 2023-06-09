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

type IProductService interface {
	// public
	GetPrdoucts(ctx context.Context, userId string, payload *dto.QueryProductsDto) (*dto.QueryProductsResponseDto, error)
	GetProductById(ctx context.Context, productId string) (*ent.Product, error)
	// private
	CreateProduct(ctx context.Context, userId string, payload *dto.CreateProductDto) (*ent.Product, error)
	UpdateProductById(ctx context.Context, userId string, productId string, payload *dto.UpdateProductDto) (*ent.Product, error)
	SoftDeleteProductById(ctx context.Context, userId string, productId string) (*ent.Product, error)
}
type ProductService struct {
	logger      *zap.Logger
	client      *ent.Client
	userRepo    repository.IUserRepository
	productRepo repository.IProductRepository
}

func NewProductService(logger *zap.Logger, client *ent.Client,
	userRepo repository.IUserRepository, productRepo repository.IProductRepository) IProductService {
	return &ProductService{
		logger:      logger,
		client:      client,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

// CreateProduct
func (productSvc *ProductService) CreateProduct(
	ctx context.Context, userId string, payload *dto.CreateProductDto) (*ent.Product, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		productSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	err = payload.Validate()
	if err != nil {
		productSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to createProduct with tx
	var result *ent.Product
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		mapped := payload.MapToSchema(constants.ProductStatus.Initiated)
		rs, err := productSvc.productRepo.CreateProduct(ctx, txc, userId, mapped)
		if err != nil {
			return err
		}
		result = rs
		return nil
	}
	err = productSvc.productRepo.WithTx(ctx, productSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPrdoucts
func (productSvc *ProductService) GetPrdoucts(
	ctx context.Context, userId string, payload *dto.QueryProductsDto) (*dto.QueryProductsResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		productSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	ensuredPaging := payload.Ensure()
	ensuredPayload := dto.NewQueryProductsDto(*ensuredPaging)

	// call repo to GetProducts
	result, err := productSvc.productRepo.GetProducts(ctx, productSvc.client, userId, ensuredPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProductById
func (productSvc *ProductService) GetProductById(
	ctx context.Context, productId string) (*ent.Product, error) {
	// validate
	_, err := uuid.Parse(productId)
	if err != nil {
		productSvc.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to GetProducts
	result, err := productSvc.productRepo.GetProductById(ctx, productSvc.client, productId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateProductById
func (productSvc *ProductService) UpdateProductById(
	ctx context.Context, userId string, productId string, payload *dto.UpdateProductDto) (*ent.Product, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		productSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	_, err = uuid.Parse(productId)
	if err != nil {
		productSvc.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	err = payload.Validate()
	if err != nil {
		productSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to updateProduct with tx
	var result *ent.Product
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		product, err := productSvc.productRepo.GetProductById(ctx, txc, productId)
		if err != nil {
			return err
		}

		if product.UserID.String() != userId {
			return constants.ErrUnauthorized
		}

		rs, err := productSvc.productRepo.UpdateProductById(ctx, txc, productId, payload)
		if err != nil {
			return err
		}
		result = rs
		return nil
	}
	err = productSvc.productRepo.WithTx(ctx, productSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SoftDeleteProductById
func (productSvc *ProductService) SoftDeleteProductById(
	ctx context.Context, userId string, productId string) (*ent.Product, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		productSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	_, err = uuid.Parse(productId)
	if err != nil {
		productSvc.logger.Info("fail to parse productId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to updateProduct with tx
	var result *ent.Product
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		product, err := productSvc.productRepo.GetProductById(ctx, txc, productId)
		if err != nil {
			return err
		}

		// check product belong to user
		if product.UserID.String() != userId {
			return constants.ErrUnauthorized
		}

		rs, err := productSvc.productRepo.SoftDeleteProductById(ctx, txc, productId)
		if err != nil {
			return err
		}
		result = rs
		return nil
	}
	err = productSvc.productRepo.WithTx(ctx, productSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}
