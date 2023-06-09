package dto

import (
	"sthl/ent"
	"sthl/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ****CreateProductDto
type CreateProductDto struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Quantity    *int32   `json:"quantity"`
	Description *string  `json:"description"`
	ImgUrl      *string  `json:"imgUrl"`
}
type CreateProductDtoMappedDto struct {
	Name        *string
	Price       *float64
	Quantity    *int32
	Description *string
	ImgUrl      *string
	Status      *string
}

func NewCreateProductDto(name *string, price *float64, quantity *int32, desc *string, imgUrl *string) *CreateProductDto {
	return &CreateProductDto{
		Name:        name,
		Price:       price,
		Quantity:    quantity,
		Description: desc,
		ImgUrl:      imgUrl,
	}
}

func (d CreateProductDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, ProductNameRule...),
		validation.Field(&d.Price, ProductPriceRule...),
		validation.Field(&d.Quantity, ProductQuantityRule...),
		validation.Field(&d.Description, ProductDescriptionRule...),
		validation.Field(&d.ImgUrl, ProductImgUrlRule...),
	)
}

func (d *CreateProductDto) MapToSchema(status string) *CreateProductDtoMappedDto {
	return &CreateProductDtoMappedDto{
		Name:        d.Name,
		Price:       d.Price,
		Quantity:    d.Quantity,
		Description: d.Description,
		ImgUrl:      d.ImgUrl,
		Status:      utils.PtrOf(status),
	}
}

// ****QueryProductsDto
type QueryProductsDto struct {
	Paging
}

func NewQueryProductsDto(paging Paging) *QueryProductsDto {
	return &QueryProductsDto{
		Paging: paging,
	}
}

type QueryProductsResponseDto struct {
	Data           []*ent.Product `json:"products"`
	PagingResponse `json:""`
}

func NewQueryProductsResponseDto(data []*ent.Product, paging PagingResponse) *QueryProductsResponseDto {
	return &QueryProductsResponseDto{
		Data:           data,
		PagingResponse: paging,
	}
}

// ****UpdateProductDto
type UpdateProductDto struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Quantity    *int32   `json:"quantity"`
	Description *string  `json:"description"`
	Status      *string  `json:"status"`
	ImgUrl      *string  `json:"imgUrl"`
}

func NewUpdateProductDto(name *string, price *float64, quantity *int32, desc *string, status *string, imgUrl *string) *UpdateProductDto {
	return &UpdateProductDto{
		Name:        name,
		Price:       price,
		Quantity:    quantity,
		Description: desc,
		Status:      status,
		ImgUrl:      imgUrl,
	}
}
func (d UpdateProductDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, ProductNameRule...),
		validation.Field(&d.Price, ProductPriceRule...),
		validation.Field(&d.Quantity, ProductQuantityRule...),
		validation.Field(&d.Description, ProductDescriptionRule...),
		validation.Field(&d.Status, ProductStatusRule...),
		validation.Field(&d.ImgUrl, ProductImgUrlRule...),
	)
}
