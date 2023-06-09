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
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// orderServiceTestSetup
func orderServiceTestSetup(ctx context.Context, t *testing.T) (*assert.Assertions, IOrderService, string, *ent.Product) {
	t.Setenv("JWT_SECRET", "test_value")
	assert := assert.New(t)
	// dependency init
	zapLogger, err := logger.NewDevInfoZapLogger()
	// zapLogger, err := logger.NewDevErrorZapLogger()
	assert.NotEmpty(zapLogger)
	assert.NoError(err)
	userRepo := repository.NewUserRepositoryMock()
	productRepo := repository.NewProductRepositoryMock()
	orderRepo := repository.NewOrderRepositoryMock()
	orderSvc := NewOrderService(zapLogger, nil, userRepo, productRepo, orderRepo)
	assert.NotEmpty(userRepo)
	assert.NotEmpty(productRepo)
	assert.NotEmpty(orderRepo)
	assert.NotEmpty(orderSvc)

	// pre
	validUserId := uuid.NewString()
	validCreateProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(100000, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))
	p1, err := productRepo.CreateProduct(ctx, nil, validUserId, validCreateProductDto.MapToSchema(constants.ProductStatus.Active))
	assert.NotEmpty(validUserId)
	assert.NotEmpty(validCreateProductDto)
	assert.NotEmpty(p1)
	assert.NoError(err)

	return assert, orderSvc, validUserId, p1
}

// pre create order
func preCreateOrder(
	ctx context.Context, assert *assert.Assertions, orderSvc IOrderService, validUserId string, product *ent.Product) *dto.OrderResponseDto {
	// pre create order
	validCreateOrderItems := dto.NewOrderItem(
		utils.PtrOf(product.ID.String()),
		utils.PtrOf(product.Name),
		utils.PtrOf(product.Price),
		utils.PtrOf(2),
	)
	validCreateOrderDto := dto.NewCreateOrderDto(
		[]*dto.OrderItem{validCreateOrderItems},
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(gofakeit.Address().Address),
	)
	preOrder, err := orderSvc.CreateOrder(ctx, validUserId, validCreateOrderDto)
	assert.NotEmpty(preOrder)
	assert.NoError(err)

	return preOrder
}

// ****

// ****Test_CreateOrder
type createOrderTestCase struct {
	name   string
	userId string
	input  *dto.CreateOrderDto
	exec   func(*dto.OrderResponseDto, error)
}

func Test_CreateOrder(t *testing.T) {
	ctx := context.TODO()
	assert, orderSvc, validUserId, p1 := orderServiceTestSetup(ctx, t)
	nonExistUserId := uuid.NewString()
	validCreateOrderItems := dto.NewOrderItem(
		utils.PtrOf(p1.ID.String()),
		utils.PtrOf(p1.Name),
		utils.PtrOf(p1.Price),
		utils.PtrOf(2),
	)
	nonExistCreateOrderItem := dto.NewOrderItem(
		utils.PtrOf(uuid.NewString()),
		utils.PtrOf(p1.Name),
		utils.PtrOf(p1.Price),
		utils.PtrOf(2),
	)
	validCreateOrderDto := dto.NewCreateOrderDto(
		[]*dto.OrderItem{validCreateOrderItems},
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(gofakeit.Address().Address),
	)
	invalidCreateOrderDto := dto.NewCreateOrderDto(
		[]*dto.OrderItem{nonExistCreateOrderItem},
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(gofakeit.Address().Address),
	)

	testCases := []createOrderTestCase{
		{
			name:   "create with valid param",
			userId: validUserId,
			input:  validCreateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "create with invalid param, nonExistUserId",
			userId: nonExistUserId,
			input:  validCreateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "create with invalid param, nonExistOrderItem",
			userId: nonExistUserId,
			input:  invalidCreateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(orderSvc.CreateOrder(ctx, test.userId, test.input))
		})
	}
}

// ****Test_GetOrders
type getOrdersTestCase struct {
	name   string
	userId string
	input  *dto.QueryOrdersDto
	exec   func(*dto.QueryOrdersResponseDto, error)
}

func Test_GetOrders(t *testing.T) {
	ctx := context.TODO()
	assert, orderSvc, validUserId, p1 := orderServiceTestSetup(ctx, t)
	nonExistUserId := uuid.NewString()

	_ = preCreateOrder(ctx, assert, orderSvc, validUserId, p1)

	testCases := []getOrdersTestCase{
		{
			name:   "invalid paging",
			userId: validUserId,
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(-1, 2, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "invalid paging, limit exceed 1000",
			userId: validUserId,
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(1, 1001, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "valid paging",
			userId: validUserId,
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(1, 10, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "valid paging, limit",
			userId: validUserId,
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "non existed userId",
			userId: nonExistUserId,
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "incorrect userId",
			userId: uuid.NewString(),
			input:  dto.NewQueryOrdersDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryOrdersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(orderSvc.GetOrders(ctx, test.userId, test.input))
		})
	}
}

// ****Test_GetOrderById
type getOrderByIdTestCase struct {
	name    string
	userId  string
	orderId string
	exec    func(*dto.OrderResponseDto, error)
}

func Test_GetOrderById(t *testing.T) {
	ctx := context.TODO()
	assert, orderSvc, validUserId, p1 := orderServiceTestSetup(ctx, t)
	nonExistUserId := uuid.NewString()
	nonExistOrderId := uuid.NewString()

	preOrder1 := preCreateOrder(ctx, assert, orderSvc, validUserId, p1)

	testCases := []getOrderByIdTestCase{
		{
			name:    "get with valid params",
			userId:  validUserId,
			orderId: preOrder1.ID.String(),
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:    "get with invalid params, nonExistUserId",
			userId:  nonExistUserId,
			orderId: preOrder1.ID.String(),
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:    "get with invalid params, nonExistOrderId",
			userId:  validUserId,
			orderId: nonExistOrderId,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(orderSvc.GetOrderById(ctx, test.userId, test.orderId))
		})
	}
}

// ****Test_UpdateOrderById
type updateOrderByIdTestCase struct {
	name    string
	userId  string
	orderId string
	input   *dto.UpdateOrderDto
	exec    func(*dto.OrderResponseDto, error)
}

func Test_UpdateOrderById(t *testing.T) {
	ctx := context.TODO()
	assert, orderSvc, validUserId, p1 := orderServiceTestSetup(ctx, t)
	nonExistUserId := uuid.NewString()
	nonExistOrderId := uuid.NewString()

	// pre create order
	preOrder1 := preCreateOrder(ctx, assert, orderSvc, validUserId, p1)
	validUpdateOrderDto := dto.NewUpdateOrderDto(
		lo.Map(preOrder1.Items, func(item *ent.OrderItem, _ int) *dto.OrderItem {
			id := item.ID.String()
			return dto.NewOrderItem(&id, &item.PurchasedName, &item.PurchasedPrice, &item.Quantity)
		}),
		&preOrder1.Remark,
		&preOrder1.Discount,
		&preOrder1.TotalAmount,
		&preOrder1.Status,
		&preOrder1.PaymentStatus,
		&preOrder1.PaymentMethod,
		&preOrder1.DeliveryStatus,
		&preOrder1.ShippingAddress,
		&preOrder1.TrackingNumber,
	)

	testCases := []updateOrderByIdTestCase{
		{
			name:    "update with valid params",
			userId:  validUserId,
			orderId: preOrder1.ID.String(),
			input:   validUpdateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:    "update with invalid params, nonExistUserId",
			userId:  nonExistUserId,
			orderId: preOrder1.ID.String(),
			input:   validUpdateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:    "update with invalid params, nonExistUserId",
			userId:  validUserId,
			orderId: nonExistOrderId,
			input:   validUpdateOrderDto,
			exec: func(result *dto.OrderResponseDto, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(orderSvc.UpdateOrderById(ctx, test.userId, test.orderId, test.input))
		})
	}
}

// ****Test_SoftDeleteOrderById
type softDeleteOrderByIdTestCase struct {
	name    string
	userId  string
	orderId string
	exec    func(bool, error)
}

func Test_SoftDeleteOrderById(t *testing.T) {
	ctx := context.TODO()
	assert, orderSvc, validUserId, p1 := orderServiceTestSetup(ctx, t)
	nonExistUserId := uuid.NewString()
	nonExistOrderId := uuid.NewString()

	// pre create order
	preOrder1 := preCreateOrder(ctx, assert, orderSvc, validUserId, p1)

	testCases := []softDeleteOrderByIdTestCase{
		{
			name:    "soft delete with invalid params, nonExistUserId",
			userId:  nonExistUserId,
			orderId: preOrder1.ID.String(),
			exec: func(result bool, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:    "soft delete with invalid params, nonExistOrderId",
			userId:  validUserId,
			orderId: nonExistOrderId,
			exec: func(result bool, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:    "soft delete with valid params",
			userId:  validUserId,
			orderId: preOrder1.ID.String(),
			exec: func(result bool, e error) {
				assert.True(result)
				assert.NoError(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(orderSvc.SoftDeleteOrderById(ctx, test.userId, test.orderId))
		})
	}
}
