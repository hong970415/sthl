package service

import (
	"context"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type IOrderService interface {
	// private
	CreateOrder(ctx context.Context, userId string, payload *dto.CreateOrderDto) (*dto.OrderResponseDto, error)
	GetOrders(ctx context.Context, userId string, payload *dto.QueryOrdersDto) (*dto.QueryOrdersResponseDto, error)
	GetOrderById(ctx context.Context, userId string, orderId string) (*dto.OrderResponseDto, error)
	UpdateOrderById(ctx context.Context, userId string, orderId string, payload *dto.UpdateOrderDto) (*dto.OrderResponseDto, error)
	SoftDeleteOrderById(ctx context.Context, userId string, orderId string) (bool, error)
}
type OrderService struct {
	logger      *zap.Logger
	client      *ent.Client
	userRepo    repository.IUserRepository
	productRepo repository.IProductRepository
	orderRepo   repository.IOrderRepository
}

func NewOrderService(logger *zap.Logger, client *ent.Client,
	userRepo repository.IUserRepository, productRepo repository.IProductRepository, orderRepo repository.IOrderRepository) IOrderService {
	return &OrderService{
		logger:      logger,
		client:      client,
		userRepo:    userRepo,
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

// CreateOrder
func (orderSvc *OrderService) CreateOrder(
	ctx context.Context, userId string, payload *dto.CreateOrderDto) (*dto.OrderResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		orderSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	err = payload.Validate()
	if err != nil {
		orderSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to createOrder with tx
	var result *dto.OrderResponseDto
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to get all related products and check
		for _, orderitem := range payload.Items {
			product, err := orderSvc.productRepo.GetProductById(ctx, txc, *orderitem.ProductId)
			if err != nil {
				return err
			}

			// check product belong user
			if product.UserID.String() != userId {
				return constants.ErrUnauthorized
			}

			// check has sufficient quantity
			if product.Quantity == 0 {
				orderSvc.logger.Info("has no sufficient quantity to be purchase")
				return constants.ErrBadRequest
			}
			// Check item.PaidPrice is equal to productDoc.Price
			if product.Price != *orderitem.PurchasedPrice {
				orderSvc.logger.Info("product price not equal item price")
				return constants.ErrBadRequest
			}
			// Check item.Quantity is not valid
			if product.Quantity < int32(*orderitem.Quantity) {
				orderSvc.logger.Info("item quantity larfe than quantity")
				return constants.ErrBadRequest
			}

			product.Quantity -= int32(*orderitem.Quantity)
			updateProductPayload := dto.NewUpdateProductDto(
				&product.Name, &product.Price, &product.Quantity, &product.Description, &product.Status, &product.ImgURL)
			_, err = orderSvc.productRepo.UpdateProductById(ctx, txc, product.ID.String(), updateProductPayload)
			if err != nil {
				return err
			}
		}

		// call repo to create order row
		mapped := payload.MapToSchema(
			constants.OrderStatus.Initiated, constants.PaymentStatus.Pending, constants.DeliveryStatus.Pending, uuid.NewString())
		rsOrder, err := orderSvc.orderRepo.CreateOrder(ctx, txc, userId, mapped)
		if err != nil {
			return err
		}

		// call repo to create order item rows
		rsOrderItems, err := orderSvc.orderRepo.CreateOrderItems(ctx, txc, rsOrder.ID.String(), payload.Items)
		if err != nil {
			return err
		}
		result = dto.NewOrderResponseDto(rsOrder, rsOrderItems)
		return nil
	}
	err = orderSvc.orderRepo.WithTx(ctx, orderSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOrders
func (orderSvc *OrderService) GetOrders(
	ctx context.Context, userId string, payload *dto.QueryOrdersDto) (*dto.QueryOrdersResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		orderSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	ensuredPaging := payload.Ensure()
	ensuredPayload := dto.NewQueryOrdersDto(*ensuredPaging)

	// call repo to getOrders with tx
	var result *dto.QueryOrdersResponseDto
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		rs, err := orderSvc.orderRepo.GetOrders(ctx, txc, userId, ensuredPayload)
		if err != nil {
			return err
		}
		result = rs
		return nil
	}

	err = orderSvc.orderRepo.WithTx(ctx, orderSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOrderById
func (orderSvc *OrderService) GetOrderById(
	ctx context.Context, userId string, orderId string) (*dto.OrderResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		orderSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	_, err = uuid.Parse(orderId)
	if err != nil {
		orderSvc.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to getOrderById with tx
	var result *dto.OrderResponseDto
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		rs, err := orderSvc.orderRepo.GetOrderById(ctx, txc, orderId)
		if err != nil {
			return err
		}

		if rs.UserID.String() != userId {
			return constants.ErrUnauthorized
		}
		result = rs
		return nil
	}
	err = orderSvc.orderRepo.WithTx(ctx, orderSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateOrderById
func (orderSvc *OrderService) UpdateOrderById(
	ctx context.Context, userId string, orderId string, payload *dto.UpdateOrderDto) (*dto.OrderResponseDto, error) {
	// validate
	_, err := uuid.Parse(userId)
	if err != nil {
		orderSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	_, err = uuid.Parse(orderId)
	if err != nil {
		orderSvc.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}
	err = payload.Validate()
	if err != nil {
		orderSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	var result *dto.OrderResponseDto
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		originalOrder, err := orderSvc.orderRepo.GetOrderById(ctx, txc, orderId)
		if err != nil {
			return err
		}

		// **handle update order item
		originalOrderItemProductIds := lo.Map(originalOrder.Items, func(item *ent.OrderItem, _ int) string { return item.ProductID.String() })
		newsetOrderItemProductIds := lo.Map(payload.Items, func(item *dto.OrderItem, _ int) string { return *item.ProductId })
		for _, originalOrderItem := range originalOrder.Items {
			originalOrderItemId := originalOrderItem.ID.String()
			originalOrderItemProductId := originalOrderItem.ProductID.String()
			// call repo to get product by id
			orderitemProduct, err := orderSvc.productRepo.GetProductById(ctx, txc, originalOrderItemProductId)
			if err != nil {
				return err
			}

			isDelete := !lo.Contains(newsetOrderItemProductIds, originalOrderItemProductId)
			// delete case
			if isDelete {
				// call repo to delete order item by id
				rsDelete, err := orderSvc.orderRepo.DeleteOrderItemById(ctx, txc, originalOrderItemId)
				if err != nil || !rsDelete {
					return err
				}

				// call repo to update back quantity for product
				updateProductPayload := dto.NewUpdateProductDto(
					&orderitemProduct.Name, &orderitemProduct.Price, &orderitemProduct.Quantity, &orderitemProduct.Description, &orderitemProduct.Status,
					&orderitemProduct.ImgURL,
				)
				*updateProductPayload.Quantity += int32(originalOrderItem.Quantity)
				_, err = orderSvc.productRepo.UpdateProductById(ctx, txc, originalOrderItemProductId, updateProductPayload)
				if err != nil {
					return err
				}
			} else { // update case
				//  find  by orderitemProductId
				orderItemFromPayload, ok := lo.Find(payload.Items, func(item *dto.OrderItem) bool { return *item.ProductId == originalOrderItemProductId })
				if !ok {
					return constants.ErrBadRequest
				}

				// create NewOrderItem and validate
				updateOrderItemPayload := dto.NewOrderItem(
					&originalOrderItemId, &originalOrderItem.PurchasedName, &originalOrderItem.PurchasedPrice, orderItemFromPayload.Quantity,
				)
				err := updateOrderItemPayload.Validate()
				if err != nil {
					return err
				}

				// new quantity - old quantity
				quantityDiff := *orderItemFromPayload.Quantity - originalOrderItem.Quantity

				// check has sufficient quantity
				if orderitemProduct.Quantity == 0 && quantityDiff > 0 {
					orderSvc.logger.Info("has no sufficient quantity to be purchase")
					return constants.ErrBadRequest
				}
				// check item.Quantity is not valid
				if orderitemProduct.Quantity < int32(quantityDiff) {
					orderSvc.logger.Info("item quantity larfe than quantity")
					return constants.ErrBadRequest
				}

				// call repo to updateOrderItemById
				_, err = orderSvc.orderRepo.UpdateOrderItemById(ctx, txc, originalOrderItemId, updateOrderItemPayload)
				if err != nil {
					return err
				}

				// call repo to update back quantity for product
				updateProductPayload := dto.NewUpdateProductDto(
					&orderitemProduct.Name, &orderitemProduct.Price, &orderitemProduct.Quantity, &orderitemProduct.Description, &orderitemProduct.Status,
					&orderitemProduct.ImgURL,
				)
				*updateProductPayload.Quantity = *updateProductPayload.Quantity - int32(quantityDiff)
				_, err = orderSvc.productRepo.UpdateProductById(ctx, txc, originalOrderItemProductId, updateProductPayload)
				if err != nil {
					return err
				}
			}
		}

		// create case
		orderitemsToCreate := lo.FilterMap(payload.Items, func(item *dto.OrderItem, _ int) (*dto.OrderItem, bool) {
			isCreate := !lo.Contains(originalOrderItemProductIds, *item.ProductId)
			if isCreate {
				return item, true
			}
			return nil, false
		})
		// call repo to create order items
		_, err = orderSvc.orderRepo.CreateOrderItems(ctx, txc, orderId, orderitemsToCreate)
		if err != nil {
			return err
		}

		// **handle update order
		_, err = orderSvc.orderRepo.UpdateOrderById(ctx, txc, orderId, payload)
		if err != nil {
			return err
		}

		// call repo to get order
		orderResp, err := orderSvc.orderRepo.GetOrderById(ctx, txc, orderId)
		if err != nil {
			return err
		}
		result = orderResp
		return nil
	}
	err = orderSvc.orderRepo.WithTx(ctx, orderSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SoftDeleteOrderById
func (orderSvc *OrderService) SoftDeleteOrderById(
	ctx context.Context, userId string, orderId string) (bool, error) {
	_, err := uuid.Parse(userId)
	if err != nil {
		orderSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}
	_, err = uuid.Parse(orderId)
	if err != nil {
		orderSvc.logger.Info("fail to parse orderId to uuid", zap.Error(err))
		return false, constants.ErrBadRequest
	}

	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to get original order
		originalOrder, err := orderSvc.orderRepo.GetOrderById(ctx, txc, orderId)
		if err != nil {
			return err
		}

		// checking
		if originalOrder.UserID.String() != userId {
			return constants.ErrUnauthorized
		}
		if originalOrder.IsArchived {
			return constants.ErrBadRequest
		}

		// call repo to soft delete order
		_, err = orderSvc.orderRepo.SoftDeleteOrderById(ctx, txc, orderId)
		if err != nil {
			return err
		}
		return nil
	}
	err = orderSvc.orderRepo.WithTx(ctx, orderSvc.client, txFunc)
	if err != nil {
		return false, err
	}
	return true, nil
}
