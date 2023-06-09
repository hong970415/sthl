package dto

import (
	"sthl/constants"
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

// ****Test_CreateProductDtoValidate
type createProductDtoValidateTestCase struct {
	name  string
	input *CreateProductDto
	exec  func(error)
}

func Test_CreateProductDtoValidate(t *testing.T) {
	assert := assert.New(t)

	validParam := NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	testCases := []createProductDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validParam,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, name empty",
			input: NewCreateProductDto(
				utils.PtrOf(""),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(501)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, description exceed limit",
			input: NewCreateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(513)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, price min",
			input: NewCreateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(float64(0.9999)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with valid param, quantity zero",
			input: NewCreateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(0)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, quantity smaller than zero",
			input: NewCreateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(-1)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, all nil",
			input: NewCreateProductDto(nil, nil, nil, nil, nil),
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

// ****Test_UpdateProductDtoValidate
type updateProductDtoValidateTestCase struct {
	name  string
	input *UpdateProductDto
	exec  func(error)
}

func Test_UpdateProductDtoValidate(t *testing.T) {
	assert := assert.New(t)

	testCases := []updateProductDtoValidateTestCase{
		{
			name: "validate with valid param",
			input: NewUpdateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, name empty",
			input: NewUpdateProductDto(
				utils.PtrOf(""),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(501)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, description exceed limit",
			input: NewUpdateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(513)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with invalid param, price min",
			input: NewUpdateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(float64(0.9999)),
				utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name: "validate with valid param, quantity zero",
			input: NewUpdateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(0)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name: "validate with invalid param, quantity smaller than zero",
			input: NewUpdateProductDto(
				utils.PtrOf(gofakeit.Fruit()),
				utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
				utils.PtrOf(int32(-1)),
				utils.PtrOf(gofakeit.LetterN(100)),
				utils.PtrOf(constants.ProductStatus.Active),
				utils.PtrOf(gofakeit.LetterN(100))),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, all nil",
			input: NewUpdateProductDto(nil, nil, nil, nil, nil, nil),
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
