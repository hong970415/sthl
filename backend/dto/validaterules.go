package dto

import (
	"errors"
	"sthl/constants"
	"sthl/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// ****User
var (
	UserEmailRule = []validation.Rule{
		validation.Required, is.EmailFormat,
	}
	UserEmailVerifiedRule = []validation.Rule{
		validation.NotNil,
	}
	UserHashedPwRule = []validation.Rule{
		validation.Required, validation.Length(6, 255),
	}
	UserPasswordRule = []validation.Rule{
		validation.Required, validation.Length(6, 255),
	}
)

// ****Product
var (
	ProductNameRule = []validation.Rule{
		validation.Required, validation.Length(2, 128),
	}
	ProductPriceRule = []validation.Rule{
		validation.Required, validation.Min(1.0),
	}
	ProductQuantityRule = []validation.Rule{
		validation.NotNil, validation.Min(0),
	}
	ProductDescriptionRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 512),
	}
	ProductStatusRule = []validation.Rule{
		validation.NotNil, validation.By(InStrings(constants.ProductStatus.GetList(), "status")),
	}
	ProductImgUrlRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 512),
	}
)

// ****Order
var (
	OrderRemarkRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 255),
	}
	OrderDiscountRule = []validation.Rule{
		validation.Required, validation.Min(0.1), validation.Max(1.0),
	}
	OrderTotalAmountRule = []validation.Rule{
		validation.NotNil, validation.Min(0.0),
	}
	OrderStatusRule = []validation.Rule{
		validation.Required, validation.By(InStrings(constants.OrderStatus.GetList(), "order status")),
	}
	OrderPaymentStatusRule = []validation.Rule{
		validation.Required, validation.By(InStrings(constants.PaymentStatus.GetList(), "payment status")),
	}
	OrderPaymentMethodRule = []validation.Rule{
		validation.Required, validation.By(InStrings(constants.PaymentMethod.GetList(), "payment method")),
	}
	OrderDeliveryStatusRule = []validation.Rule{
		validation.Required, validation.By(InStrings(constants.DeliveryStatus.GetList(), "delivery status")),
	}
	OrderShippingAddressRule = []validation.Rule{
		validation.Required, validation.Length(1, 255),
	}
	OrderTrackingNumberRule = []validation.Rule{
		validation.Required, validation.Length(1, 128),
	}
	checkOrderItemsProductIdIsUnique = func(value interface{}) error {
		s, ok := value.([]*OrderItem)
		if !ok {
			return errors.New("fail to parse value to []*OrderItem")
		}
		ids := []string{}
		for _, item := range s {
			err := item.Validate()
			if err != nil {
				return err
			}
			_, foundIndex := utils.Find(ids, func(index int, id string) bool { return id == *item.ProductId })
			if foundIndex != -1 {
				return errors.New("order items not unique")
			}
			ids = append(ids, *item.ProductId)
		}
		return nil
	}
	OrderItemsRule = []validation.Rule{
		validation.Required, validation.By(checkOrderItemsProductIdIsUnique),
	}
	OrderItemIdRule = []validation.Rule{
		validation.Required, is.UUID, validation.By(NotEquals(utils.PtrOf(uuid.Nil.String()), "order item id and zero uuid")),
	}
	OrderItemProductIdRule = []validation.Rule{
		validation.Required, is.UUID, validation.By(NotEquals(utils.PtrOf(uuid.Nil.String()), "order item productId and zero uuid")),
	}
	OrderItemPurchasedNameRule = []validation.Rule{
		validation.Required, validation.Length(2, 128),
	}
	OrderItemPurchasedPriceRule = []validation.Rule{
		validation.Required, validation.Min(1.0),
	}
	OrderItemQuantityRule = []validation.Rule{
		validation.Required, validation.Min(1),
	}
	// SiteUi
	SiteNameRule = []validation.Rule{
		validation.Required, validation.Length(1, 32),
	}
	HomepageImgUrlRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 512),
	}
	HomepageTextRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 255),
	}
	HomepageTextColorRule = []validation.Rule{
		validation.NotNil, validation.Length(0, 16),
	}
)
