package dto

import (
	"sthl/ent"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ****CreateOrderDto

// OrderItem
type OrderItem struct {
	ProductId      *string  `json:"productId"`
	PurchasedName  *string  `json:"purchasedName"`
	PurchasedPrice *float64 `json:"purchasedPrice"`
	Quantity       *int     `json:"quantity"`
}

func NewOrderItem(productId *string, purchasedName *string, purchasedPrice *float64, quantity *int) *OrderItem {
	return &OrderItem{
		ProductId:      productId,
		PurchasedName:  purchasedName,
		PurchasedPrice: purchasedPrice,
		Quantity:       quantity,
	}
}
func (d OrderItem) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.ProductId, OrderItemProductIdRule...),
		validation.Field(&d.PurchasedName, OrderItemPurchasedNameRule...),
		validation.Field(&d.PurchasedPrice, OrderItemPurchasedPriceRule...),
		validation.Field(&d.Quantity, OrderItemQuantityRule...),
	)
}

// CreateOrderDto
type CreateOrderDto struct {
	Items           []*OrderItem `json:"items"`
	Remark          *string      `json:"remark"`
	Discount        *float64     `json:"discount"`
	TotalAmount     *float64     `json:"totalAmount"`
	PaymentMethod   *string      `json:"paymentMethod"`
	ShippingAddress *string      `json:"shippingAddress"`
}

func NewCreateOrderDto(
	items []*OrderItem, remark *string, discount *float64, totalAmount *float64, paymentMethod *string, shippingAddress *string) *CreateOrderDto {
	return &CreateOrderDto{
		Items:           items,
		Remark:          remark,
		Discount:        discount,
		TotalAmount:     totalAmount,
		PaymentMethod:   paymentMethod,
		ShippingAddress: shippingAddress,
	}
}
func (d CreateOrderDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Items, OrderItemsRule...),
		validation.Field(&d.Remark, OrderRemarkRule...),
		validation.Field(&d.Discount, OrderDiscountRule...),
		validation.Field(&d.TotalAmount, OrderTotalAmountRule...),
		validation.Field(&d.PaymentMethod, OrderPaymentMethodRule...),
		validation.Field(&d.ShippingAddress, OrderShippingAddressRule...),
	)
}

// CreateOrderDtoMappedDto
type CreateOrderDtoMappedDto struct {
	Items           []*OrderItem
	Remark          *string
	Discount        *float64
	TotalAmount     *float64
	Status          *string
	PaymentStatus   *string
	PaymentMethod   *string
	DeliveryStatus  *string
	ShippingAddress *string
	TrackingNumber  *string
}

func (d *CreateOrderDto) MapToSchema(status string, paymentStatus string, deliveryStatus string, trackingNumber string) *CreateOrderDtoMappedDto {
	return &CreateOrderDtoMappedDto{
		Items:           d.Items,
		Remark:          d.Remark,
		Discount:        d.Discount,
		TotalAmount:     d.TotalAmount,
		Status:          &status,
		PaymentStatus:   &paymentStatus,
		PaymentMethod:   d.PaymentMethod,
		DeliveryStatus:  &deliveryStatus,
		ShippingAddress: d.ShippingAddress,
		TrackingNumber:  &trackingNumber,
	}
}

type OrderResponseDto struct {
	*ent.Order `json:","`
	Items      []*ent.OrderItem `json:"items"`
}

func NewOrderResponseDto(order *ent.Order, orderItems []*ent.OrderItem) *OrderResponseDto {
	return &OrderResponseDto{
		order,
		orderItems,
	}
}

// ****QueryOrdersDto
type QueryOrdersDto struct {
	Paging
}

func NewQueryOrdersDto(paging Paging) *QueryOrdersDto {
	return &QueryOrdersDto{
		Paging: paging,
	}
}

type QueryOrdersResponseDto struct {
	Data           []*OrderResponseDto `json:"orders"`
	PagingResponse `json:""`
}

func NewQueryOrdersResponseDto(data []*OrderResponseDto, paging PagingResponse) *QueryOrdersResponseDto {
	return &QueryOrdersResponseDto{
		Data:           data,
		PagingResponse: paging,
	}
}

// ****UpdateOrderDto
type UpdateOrderDto struct {
	Items           []*OrderItem `json:"items"`
	Remark          *string      `json:"remark"`
	Discount        *float64     `json:"discount"`
	TotalAmount     *float64     `json:"totalAmount"`
	Status          *string      `json:"status"`
	PaymentStatus   *string      `json:"paymentStatus"`
	PaymentMethod   *string      `json:"paymentMethod"`
	DeliveryStatus  *string      `json:"deliveryStatus"`
	ShippingAddress *string      `json:"shippingAddress"`
	TrackingNumber  *string      `json:"trackingNumber"`
}

func NewUpdateOrderDto(
	items []*OrderItem, remark *string, discount *float64, totalAmount *float64,
	status *string, paymentStatus *string, paymentMethod *string,
	deliveryStatus *string, shippingAddress *string, trackingNumber *string) *UpdateOrderDto {
	return &UpdateOrderDto{
		Items:           items,
		Remark:          remark,
		Discount:        discount,
		TotalAmount:     totalAmount,
		Status:          status,
		PaymentStatus:   paymentStatus,
		PaymentMethod:   paymentMethod,
		DeliveryStatus:  deliveryStatus,
		ShippingAddress: shippingAddress,
		TrackingNumber:  trackingNumber,
	}
}

func (d UpdateOrderDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Items, OrderItemsRule...),
		validation.Field(&d.Remark, OrderRemarkRule...),
		validation.Field(&d.Discount, OrderDiscountRule...),
		validation.Field(&d.TotalAmount, OrderTotalAmountRule...),
		validation.Field(&d.Status, OrderStatusRule...),
		validation.Field(&d.PaymentStatus, OrderPaymentStatusRule...),
		validation.Field(&d.PaymentMethod, OrderPaymentMethodRule...),
		validation.Field(&d.DeliveryStatus, OrderDeliveryStatusRule...),
		validation.Field(&d.ShippingAddress, OrderShippingAddressRule...),
		validation.Field(&d.TrackingNumber, OrderTrackingNumberRule...),
	)
}
