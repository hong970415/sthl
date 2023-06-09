package constants

import (
	"time"
)

const (
	// Auth
	AccessTokenKey       contextKey    = "accessToken"
	AccessTokenInfoKey   contextKey    = "accessTokenInfo"
	AccessTokenDuration  time.Duration = 1200 * time.Hour
	RefreshTokenKey      contextKey    = "refreshToken"
	RefreshTokenDuration time.Duration = 2400 * time.Hour
	// DB
	AccountServiceDbName string = "account_db"
	// s3
	S3BucketName string = "sthl-dev"
	// User pw
	UserPwHashCost int = 10
	// file
	MaxFileSize  int64 = 4 << 20
	MaxProducts  int   = 1000
	MaxAlbumImgs int   = 1000
)

var (
	// Product Status
	ProductStatus = productStatusType{
		Initiated:  "initiated",
		Active:     "active",
		OutOfStock: "outOfStock",
		Inactive:   "inactive",
	}
	// Order Status
	OrderStatus = orderStatusType{
		Initiated: "initiated",
		Confirmed: "confirmed",
		Shipping:  "shipping",
		Canceled:  "canceled",
		Completed: "completed",
	}
	// Payment Status
	PaymentStatus = paymentStatusType{
		Pending:         "pending",
		Fail:            "fail",
		Paid:            "paid",
		Refunded:        "refunded",
		PartialRefunded: "partialRefunded",
		NoRefund:        "noRefund",
		Voided:          "voided",
	}
	// PaymentMethod
	PaymentMethod = paymentMethodType{
		Card: "card",
	}
	// Delivery Status
	DeliveryStatus = deliveryStatusType{
		Pending:   "pending",
		Shipping:  "shipping",
		Completed: "completed",
	}
)
