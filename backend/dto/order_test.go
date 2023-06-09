package dto

import (
	"sthl/constants"
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func generateMockOrderItems(nums uint) []*OrderItem {
	result := []*OrderItem{}
	for i := uint(0); i < nums; i++ {
		item := NewOrderItem(
			utils.PtrOf(uuid.NewString()),
			utils.PtrOf(gofakeit.Fruit()),
			utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
			utils.PtrOf(gofakeit.IntRange(0, 1000000)),
		)

		result = append(result, item)
	}
	return result
}

// func generateMockOrderItemsWithId(nums uint) []*OrderItemWithId {
// 	result := []*OrderItemWithId{}
// 	for i := uint(0); i < nums; i++ {
// 		item := NewOrderItemWithId(
// 			utils.PtrOf(uuid.NewString()),
// 			utils.PtrOf(uuid.NewString()),
// 			utils.PtrOf(gofakeit.Fruit()),
// 			utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
// 			utils.PtrOf(gofakeit.IntRange(0, 1000000)),
// 		)

// 		result = append(result, item)
// 	}
// 	return result
// }

// ****Test_CreateOrderDtoValidate
type createOrderDtoValidateTestCase struct {
	name  string
	input *CreateOrderDto
	exec  func(error)
}

func Test_CreateOrderDtoValidate(t *testing.T) {
	assert := assert.New(t)

	validOrderItems := generateMockOrderItems(1)
	validOrderItems2 := generateMockOrderItems(2)
	duplicateOrderItems := append(validOrderItems, validOrderItems...)
	validOrder := NewCreateOrderDto(
		validOrderItems,
		utils.PtrOf(gofakeit.LetterN(10)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(gofakeit.Address().Address),
	)
	validOrder2 := NewCreateOrderDto(
		validOrderItems2,
		utils.PtrOf(gofakeit.LetterN(10)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(gofakeit.Address().Address),
	)

	testCases := []createOrderDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validOrder,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:  "validate with valid param",
			input: validOrder2,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, duplicate items",
			input: NewCreateOrderDto(
				duplicateOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil items",
			input: NewCreateOrderDto(
				nil,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty items",
			input: NewCreateOrderDto(
				[]*OrderItem{},
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil remark",
			input: NewCreateOrderDto(
				validOrderItems,
				nil,
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with valid param, empty remark",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(0)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, nil discount",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				nil,
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, zero discount",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(0.0),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil paymentMethod",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				nil,
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty paymentMethod",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(""),
				utils.PtrOf(gofakeit.Address().Address),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil shippingAddress",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				nil,
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty shippingAddress",
			input: NewCreateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(""),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}

// ****Test_UpdateOrderDtoValidate
type updateOrderDtoValidateTestCase struct {
	name  string
	input *UpdateOrderDto
	exec  func(error)
}

func Test_UpdateOrderDtoValidate(t *testing.T) {
	assert := assert.New(t)

	// validOrderItems := generateMockOrderItemsWithId(1)
	validOrderItems := generateMockOrderItems(1)
	validOrder := NewUpdateOrderDto(
		validOrderItems,
		utils.PtrOf(gofakeit.LetterN(10)),
		utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
		utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
		utils.PtrOf(constants.OrderStatus.Initiated),
		utils.PtrOf(constants.PaymentStatus.Pending),
		utils.PtrOf(constants.PaymentMethod.Card),
		utils.PtrOf(constants.DeliveryStatus.Pending),
		utils.PtrOf(gofakeit.Address().Address),
		utils.PtrOf(gofakeit.UUID()),
	)

	testCases := []updateOrderDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validOrder,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, nil items",
			input: NewUpdateOrderDto(
				nil,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty items",
			input: NewUpdateOrderDto(
				[]*OrderItem{},
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil remark",
			input: NewUpdateOrderDto(
				validOrderItems,
				nil,
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with valid param, empty remark",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(0)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, nil discount",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				nil,
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, zero discount",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(0.0),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil paymentMethod",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				nil,
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty paymentMethod",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(""),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(gofakeit.Address().Address),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, nil shippingAddress",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				nil,
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, empty shippingAddress",
			input: NewUpdateOrderDto(
				validOrderItems,
				utils.PtrOf(gofakeit.LetterN(10)),
				utils.PtrOf(gofakeit.Float64Range(0.1, 1.0)),
				utils.PtrOf(gofakeit.Float64Range(0, 10000000)),
				utils.PtrOf(constants.OrderStatus.Initiated),
				utils.PtrOf(constants.PaymentStatus.Pending),
				utils.PtrOf(constants.PaymentMethod.Card),
				utils.PtrOf(constants.DeliveryStatus.Pending),
				utils.PtrOf(""),
				utils.PtrOf(gofakeit.UUID()),
			),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}
