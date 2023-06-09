package service

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/logger"
	"sthl/repository"
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// productServiceTestSetup
func productServiceTestSetup(ctx context.Context, t *testing.T) (*assert.Assertions, IProductService) {
	t.Setenv("JWT_SECRET", "test_value")
	assert := assert.New(t)
	// dependency init
	// zapLogger, err := logger.NewDevInfoZapLogger()
	zapLogger, err := logger.NewDevErrorZapLogger()
	assert.NotEmpty(zapLogger)
	assert.NoError(err)
	userRepo := repository.NewUserRepositoryMock()
	productRepo := repository.NewProductRepositoryMock()
	productService := NewProductService(zapLogger, nil, userRepo, productRepo)
	assert.NotEmpty(userRepo)
	assert.NotEmpty(productRepo)
	assert.NotEmpty(productService)
	return assert, productService
}

// ****Test_CreateProduct
type createProductTestCase struct {
	name   string
	userId string
	input  *dto.CreateProductDto
	exec   func(*ent.Product, error)
}

func Test_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	assert, productSvc := productServiceTestSetup(ctx, t)
	validUserId := uuid.NewString()
	invalidUserId := "invalid userid"
	validCreateProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	testCases := []createProductTestCase{
		{
			name:   "create with valid param",
			userId: validUserId,
			input:  validCreateProductDto,
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "create with valid info, empty desc",
			userId: validUserId,
			input: dto.NewCreateProductDto(
				utils.PtrOf(gofakeit.Vegetable()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(0)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "create with invalid info, duplicate name",
			userId: validUserId,
			input: dto.NewCreateProductDto(
				utils.PtrOf(*validCreateProductDto.Name),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "create with invalid userId",
			userId: invalidUserId,
			input: dto.NewCreateProductDto(
				utils.PtrOf(gofakeit.Breakfast()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "create with invalid price",
			userId: validUserId,
			input: dto.NewCreateProductDto(
				utils.PtrOf(gofakeit.Breakfast()),
				utils.PtrOf(0.99),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "create with valid quantity 0",
			userId: validUserId,
			input: dto.NewCreateProductDto(
				utils.PtrOf(gofakeit.Breakfast()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(0)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(productSvc.CreateProduct(ctx, test.userId, test.input))
		})
	}
}

// ****Test_GetProducts
type getProductsTestCase struct {
	name   string
	userId string
	input  *dto.QueryProductsDto
	exec   func(*dto.QueryProductsResponseDto, error)
}

func Test_GetProducts(t *testing.T) {
	ctx := context.TODO()
	assert, productSvc := productServiceTestSetup(ctx, t)

	validUserId := uuid.NewString()
	invalidUserId := validUserId + "wrong"

	validProductA := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))
	validProductB := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Animal()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	productAId, err := productSvc.CreateProduct(ctx, validUserId, validProductA)
	assert.NotEmpty(productAId)
	assert.NoError(err)

	productBId, err := productSvc.CreateProduct(ctx, validUserId, validProductB)
	assert.NotEmpty(productBId)
	assert.NoError(err)

	testCases := []getProductsTestCase{
		{
			name:   "invalid paging",
			userId: validUserId,
			input:  dto.NewQueryProductsDto(*dto.NewPaging(-1, 2, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "invalid paging, limit exceed 1000",
			userId: validUserId,
			input:  dto.NewQueryProductsDto(*dto.NewPaging(1, 1001, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "valid paging",
			userId: validUserId,
			input:  dto.NewQueryProductsDto(*dto.NewPaging(1, 10, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "valid paging, limit",
			userId: validUserId,
			input:  dto.NewQueryProductsDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "invalid userId",
			userId: invalidUserId,
			input:  dto.NewQueryProductsDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "incorrect userId",
			userId: uuid.NewString(),
			input:  dto.NewQueryProductsDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryProductsResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(productSvc.GetPrdoucts(ctx, test.userId, test.input))
		})
	}
}

// ****Test_GetProductById
type getProductByIdTestCase struct {
	name      string
	productId string
	exec      func(*ent.Product, error)
}

func Test_GetProductById(t *testing.T) {
	ctx := context.TODO()
	assert, productSvc := productServiceTestSetup(ctx, t)

	validUserId := uuid.NewString()
	invalidUserId := validUserId + "wrong"

	validProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))
	validProduct, err := productSvc.CreateProduct(ctx, validUserId, validProductDto)
	assert.NotEmpty(validProduct)
	assert.NoError(err)

	testCases := []getProductByIdTestCase{
		{
			name:      "get with correct info",
			productId: validProduct.ID.String(),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:      "get with invalid userId",
			productId: invalidUserId,
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "get with incorrect productId",
			productId: uuid.NewString(),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(productSvc.GetProductById(ctx, test.productId))
		})
	}
}

// ****Test_UpdateProductById
type updateProductByIdTestCase struct {
	name      string
	userId    string
	productId string
	input     *dto.UpdateProductDto
	exec      func(*ent.Product, error)
}

func Test_UpdateProductById(t *testing.T) {
	ctx := context.TODO()
	assert, productSvc := productServiceTestSetup(ctx, t)

	validUserId := uuid.NewString()
	invalidUserId := validUserId + "wrong"

	validProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	validProduct, err := productSvc.CreateProduct(ctx, validUserId, validProductDto)
	invalidProductId := validProduct.ID.String() + "wrong"
	assert.NotEmpty(validProduct)
	assert.NoError(err)

	testCases := []updateProductByIdTestCase{
		{
			name:      "success update",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Vegetable()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:      "update with invalid productId",
			userId:    validUserId,
			productId: invalidProductId,
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid userId",
			userId:    invalidUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, name",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf("a"),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, name nil",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				nil,
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, price",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(-1.0),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, price nil",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				nil,
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, price 0",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(0.0),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "update with invalid params, quantity nil",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				nil,
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "success update with valid params, quantity 0",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(0)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:      "success update with valid params, desc empty",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(""),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:      "update with invalid params, desc nil",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			input: dto.NewUpdateProductDto(
				utils.PtrOf(gofakeit.Lunch()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				nil,
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100)),
			),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		// {
		// 	name:      "update with invalid params, status archived",
		// 	userId:    validUserId,
		// 	productId: validProduct.ID.String(),
		// 	input: dto.NewUpdateProductDto(
		// 		utils.PtrOf(gofakeit.Lunch()),
		// 		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		// 		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		// 		utils.PtrOf(gofakeit.LetterN(100)),
		// 		utils.PtrOf(constants.ProductStatus.Archived),
		// 	),
		// 	exec: func(result *ent.Product, e error) {
		// 		assert.Empty(result)
		// 		assert.Error(e)
		// 	},
		// },
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(productSvc.UpdateProductById(ctx, test.userId, test.productId, test.input))
		})
	}
}

// ****Test_DeleteProductById
type deleteProductByIdTestCase struct {
	name      string
	userId    string
	productId string
	exec      func(*ent.Product, error)
}

func Test_DeleteProductById(t *testing.T) {
	ctx := context.TODO()
	assert, productSvc := productServiceTestSetup(ctx, t)

	validUserId := uuid.NewString()
	invalidUserId := validUserId + "wrong"

	validProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	validProduct, err := productSvc.CreateProduct(ctx, validUserId, validProductDto)
	invalidProductId := validProduct.ID.String() + "wrong"
	assert.NotEmpty(validProduct)
	assert.NoError(err)

	testCases := []deleteProductByIdTestCase{
		{
			name:      "soft delete with invalid userId",
			userId:    invalidUserId,
			productId: validProduct.ID.String(),
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "soft delete with invalid userId",
			userId:    validUserId,
			productId: invalidProductId,
			exec: func(result *ent.Product, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:      "success soft delete",
			userId:    validUserId,
			productId: validProduct.ID.String(),
			exec: func(result *ent.Product, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)

				assert.Equal(true, result.IsArchived)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(productSvc.SoftDeleteProductById(ctx, test.userId, test.productId))
		})
	}
}
