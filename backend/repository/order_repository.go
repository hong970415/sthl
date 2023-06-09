package repository

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/ent/order"
	"sthl/ent/orderitem"
	"sthl/storage"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IOrderRepository interface {
	WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error
	CreateOrder(ctx context.Context, client *ent.Client, userId string, payload *dto.CreateOrderDtoMappedDto) (*ent.Order, error)
	CreateOrderItems(ctx context.Context, client *ent.Client, orderId string, payload []*dto.OrderItem) ([]*ent.OrderItem, error)
	getOrderItemsByOrderId(ctx context.Context, client *ent.Client, orderId string) ([]*ent.OrderItem, error)
	GetOrders(ctx context.Context, client *ent.Client, userId string, payload *dto.QueryOrdersDto) (*dto.QueryOrdersResponseDto, error)
	GetOrderById(ctx context.Context, client *ent.Client, orderId string) (*dto.OrderResponseDto, error)
	UpdateOrderById(ctx context.Context, client *ent.Client, orderId string, payload *dto.UpdateOrderDto) (*ent.Order, error)
	UpdateOrderItemById(ctx context.Context, client *ent.Client, orderItemId string, payload *dto.OrderItem) (*ent.OrderItem, error)
	SoftDeleteOrderById(ctx context.Context, client *ent.Client, orderId string) (*ent.Order, error)
	DeleteOrderItemById(ctx context.Context, client *ent.Client, orderItemId string) (bool, error)
}

type OrderRepository struct {
	logger *zap.Logger
}

func NewOrderRepository(logger *zap.Logger) IOrderRepository {
	return &OrderRepository{
		logger: logger,
	}
}
func (orderRepo *OrderRepository) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTx(ctx, orderRepo.logger, client, fn)
}

// CreateOrder
func (orderRepo *OrderRepository) CreateOrder(
	ctx context.Context, client *ent.Client, userId string, payload *dto.CreateOrderDtoMappedDto) (*ent.Order, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		orderRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Order.Create().
		SetUserID(userUuid).
		SetDiscount(*payload.Discount).
		SetTotalAmount(*payload.TotalAmount).
		SetRemark(*payload.Remark).
		SetStatus(*payload.Status).
		SetPaymentStatus(*payload.PaymentStatus).
		SetPaymentMethod(*payload.PaymentMethod).
		SetDeliveryStatus(*payload.DeliveryStatus).
		SetShippingAddress(*payload.ShippingAddress).
		SetTrackingNumber(*payload.TrackingNumber).
		Save(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.Order.Create", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// CreateOrderItems
func (orderRepo *OrderRepository) CreateOrderItems(
	ctx context.Context, client *ent.Client, orderId string, payload []*dto.OrderItem) ([]*ent.OrderItem, error) {
	orderUuid, err := uuid.Parse(orderId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	bulk := make([]*ent.OrderItemCreate, len(payload))
	for i, item := range payload {
		productUuid, err := uuid.Parse(*item.ProductId)
		if err != nil {
			orderRepo.logger.Info("fail to parse item.ProductId to uuid", zap.Error(err))
			return nil, constants.ErrBadRequest
		}

		bulk[i] = client.OrderItem.Create().
			SetOrderID(orderUuid).
			SetProductID(productUuid).
			SetPurchasedName(*item.PurchasedName).
			SetPurchasedPrice(*item.PurchasedPrice).
			SetQuantity(*item.Quantity)
	}
	result, err := client.OrderItem.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.OrderItem.CreateBulk", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// getOrderItemsByOrderId
func (orderRepo *OrderRepository) getOrderItemsByOrderId(
	ctx context.Context, client *ent.Client, orderId string) ([]*ent.OrderItem, error) {
	orderUuid, err := uuid.Parse(orderId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.OrderItem.Query().
		Where(orderitem.OrderID(orderUuid)).
		All(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.OrderItem.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// GetOrders
func (orderRepo *OrderRepository) GetOrders(
	ctx context.Context, client *ent.Client, userId string, payload *dto.QueryOrdersDto) (*dto.QueryOrdersResponseDto, error) {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		orderRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit

	total, err := client.Order.Query().
		Where(order.UserID(userUuid)).Count(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to count total", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	// call ent client to Query
	rsOrders, err := client.Order.Query().
		Where(order.UserID(userUuid)).
		Order(ent.Desc(order.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.Order.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	// rsOrderItems,err
	var result []*dto.OrderResponseDto
	for _, item := range rsOrders {
		rsOrderItems, err := orderRepo.getOrderItemsByOrderId(ctx, client, item.ID.String())
		if err != nil {
			return nil, err
		}
		rs := dto.NewOrderResponseDto(item, rsOrderItems)
		result = append(result, rs)
	}

	pagingResp := dto.NewPagingResponse(page, limit, total)
	data := dto.NewQueryOrdersResponseDto(result, *pagingResp)
	return data, nil
}

// GetOrderById
func (orderRepo *OrderRepository) GetOrderById(
	ctx context.Context, client *ent.Client, orderId string) (*dto.OrderResponseDto, error) {
	orderUUid, err := uuid.Parse(orderId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	rsOrder, err := client.Order.Get(ctx, orderUUid)
	if err != nil {
		orderRepo.logger.Info("fail to client.Order.Get", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	rsOrderItems, err := orderRepo.getOrderItemsByOrderId(ctx, client, rsOrder.ID.String())
	if err != nil {
		return nil, err
	}

	result := dto.NewOrderResponseDto(rsOrder, rsOrderItems)
	return result, nil
}

// UpdateOrderById
func (orderRepo *OrderRepository) UpdateOrderById(
	ctx context.Context, client *ent.Client, orderId string, payload *dto.UpdateOrderDto) (*ent.Order, error) {
	orderUUid, err := uuid.Parse(orderId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Order.UpdateOneID(orderUUid).
		SetRemark(*payload.Remark).
		SetDiscount(*payload.Discount).
		SetTotalAmount(*payload.TotalAmount).
		SetRemark(*payload.Remark).
		SetStatus(*payload.Status).
		SetPaymentStatus(*payload.PaymentStatus).
		SetPaymentMethod(*payload.PaymentMethod).
		SetDeliveryStatus(*payload.DeliveryStatus).
		SetShippingAddress(*payload.ShippingAddress).
		SetTrackingNumber(*payload.TrackingNumber).
		Save(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.Order.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// UpdateOrderItems
func (orderRepo *OrderRepository) UpdateOrderItemById(
	ctx context.Context, client *ent.Client, orderItemId string, payload *dto.OrderItem) (*ent.OrderItem, error) {
	orderItemUUid, err := uuid.Parse(orderItemId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderItemId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	result, err := client.OrderItem.UpdateOneID(orderItemUUid).
		SetPurchasedName(*payload.PurchasedName).
		SetPurchasedPrice(*payload.PurchasedPrice).
		SetQuantity(*payload.Quantity).
		Save(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.OrderItem.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// SoftDeleteOrderById
func (orderRepo *OrderRepository) SoftDeleteOrderById(
	ctx context.Context, client *ent.Client, orderId string) (*ent.Order, error) {
	orderUUid, err := uuid.Parse(orderId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	result, err := client.Order.UpdateOneID(orderUUid).
		SetIsArchived(true).
		Save(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.OrderItem.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// DeleteOrderItemById
func (orderRepo *OrderRepository) DeleteOrderItemById(
	ctx context.Context, client *ent.Client, orderItemId string) (bool, error) {
	orderItemUUid, err := uuid.Parse(orderItemId)
	if err != nil {
		orderRepo.logger.Info("fail to parse orderItemIds to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}

	err = client.OrderItem.DeleteOneID(orderItemUUid).Exec(ctx)
	if err != nil {
		orderRepo.logger.Info("fail to client.OrderItem.DeleteOneID", zap.Error(err))
		return false, handleEntRepoErr(err)
	}
	return true, nil
}
