package repository

import (
	"context"
	"sort"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/storage"
	"sthl/utils"
	"sync"
	"time"

	"github.com/google/uuid"
)

type OrderRepositoryMock struct {
	mockDataOrder     map[string]ent.Order
	mockDataOrderItem map[string]ent.OrderItem
	mu                sync.Mutex
}

func NewOrderRepositoryMock() IOrderRepository {
	// func NewOrderRepositoryMock() any {
	return &OrderRepositoryMock{
		mockDataOrder:     map[string]ent.Order{},
		mockDataOrderItem: map[string]ent.OrderItem{},
	}
}

func newOrderSchemaMock(userId uuid.UUID, payload *dto.CreateOrderDtoMappedDto) *ent.Order {
	t := time.Now()
	return &ent.Order{
		ID:              uuid.New(),
		CreatedAt:       t,
		UpdatedAt:       t,
		UserID:          userId,
		Discount:        *payload.Discount,
		TotalAmount:     *payload.TotalAmount,
		Remark:          *payload.Remark,
		Status:          *payload.Status,
		PaymentStatus:   *payload.PaymentStatus,
		PaymentMethod:   *payload.PaymentMethod,
		DeliveryStatus:  *payload.DeliveryStatus,
		ShippingAddress: *payload.ShippingAddress,
		TrackingNumber:  *payload.TrackingNumber,
		IsArchived:      false,
	}
}
func newOrderItemSchemaMock(orderId uuid.UUID, payload *dto.OrderItem) *ent.OrderItem {
	productUuid, _ := uuid.Parse(*payload.ProductId)
	return &ent.OrderItem{
		ID:             uuid.New(),
		OrderID:        orderId,
		ProductID:      productUuid,
		PurchasedName:  *payload.PurchasedName,
		PurchasedPrice: *payload.PurchasedPrice,
		Quantity:       *payload.Quantity,
	}
}

func (m *OrderRepositoryMock) Lock() {
	m.mu.Lock()
	defer m.mu.Unlock()
}
func (m *OrderRepositoryMock) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTxTest(ctx, nil, client, fn)
}

// ****
// CreateOrder
func (m *OrderRepositoryMock) CreateOrder(ctx context.Context, client *ent.Client, userId string, payload *dto.CreateOrderDtoMappedDto) (*ent.Order, error) {
	m.Lock()
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	result := newOrderSchemaMock(userUuid, payload)
	m.mockDataOrder[result.ID.String()] = *result
	return result, nil
}

// CreateOrderItems
func (m *OrderRepositoryMock) CreateOrderItems(ctx context.Context, client *ent.Client, orderId string, payload []*dto.OrderItem) ([]*ent.OrderItem, error) {
	// return nil, constants.ErrInternalServer
	m.Lock()
	orderUuid, err := uuid.Parse(orderId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	var result []*ent.OrderItem
	for _, item := range payload {
		oi := newOrderItemSchemaMock(orderUuid, item)
		m.mockDataOrderItem[oi.ID.String()] = *oi
		result = append(result, oi)
	}
	return result, nil

}

// getOrderItemsByOrderId
func (m *OrderRepositoryMock) getOrderItemsByOrderId(ctx context.Context, client *ent.Client, orderId string) ([]*ent.OrderItem, error) {
	m.Lock()
	_, err := uuid.Parse(orderId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	var result []*ent.OrderItem
	for _, item := range m.mockDataOrderItem {
		if item.OrderID.String() == orderId {
			result = append(result, &item)
		}
	}
	return result, nil
}

// GetOrders
func (m *OrderRepositoryMock) GetOrders(ctx context.Context, client *ent.Client, userId string, payload *dto.QueryOrdersDto) (*dto.QueryOrdersResponseDto, error) {
	// return nil, constants.ErrInternalServer
	m.Lock()
	var orderSlice, subSlice []*ent.Order
	for _, u := range m.mockDataOrder {
		if u.UserID.String() == userId {
			orderSlice = append(orderSlice, &u)
		}
	}
	sort.Slice(orderSlice, func(i, j int) bool {
		return orderSlice[i].CreatedAt.Before(orderSlice[j].CreatedAt)
	})

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit
	subSlice = orderSlice[offset:]
	if len(orderSlice) > offset+limit-1 {
		subSlice = orderSlice[offset : offset+limit-1]
	}

	var data []*dto.OrderResponseDto
	for _, o := range subSlice {
		oitems, err := m.getOrderItemsByOrderId(ctx, client, o.ID.String())
		if err != nil {
			return nil, constants.ErrBadRequest
		}
		data = append(data, dto.NewOrderResponseDto(o, oitems))
	}

	pagingResp := dto.NewPagingResponse(payload.Page, payload.Limit, len(orderSlice))
	result := dto.NewQueryOrdersResponseDto(data, *pagingResp)
	return result, nil
}

// GetOrderById
func (m *OrderRepositoryMock) GetOrderById(ctx context.Context, client *ent.Client, orderId string) (*dto.OrderResponseDto, error) {
	// return nil, constants.ErrInternalServer
	m.Lock()

	_, err := uuid.Parse(orderId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	// var orderResp *dto.OrderResponseDto
	orderData := m.mockDataOrder[orderId]
	if utils.IsEmpty(orderData) {
		return nil, constants.ErrNotFound
	}
	orderItems, err := m.getOrderItemsByOrderId(ctx, client, orderId)
	if err != nil {
		return nil, err
	}
	result := dto.NewOrderResponseDto(&orderData, orderItems)

	return result, nil
}

// UpdateOrderById
func (m *OrderRepositoryMock) UpdateOrderById(ctx context.Context, client *ent.Client, orderId string, payload *dto.UpdateOrderDto) (*ent.Order, error) {
	m.Lock()

	_, err := uuid.Parse(orderId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	for key, data := range m.mockDataOrder {
		if key == orderId {
			u := data
			u.Remark = *payload.Remark
			u.Discount = *payload.Discount
			u.TotalAmount = *payload.TotalAmount
			u.Status = *payload.Status
			u.PaymentStatus = *payload.PaymentStatus
			u.PaymentMethod = *payload.PaymentMethod
			u.DeliveryStatus = *payload.DeliveryStatus
			u.ShippingAddress = *payload.ShippingAddress
			u.TrackingNumber = *payload.TrackingNumber

			m.mockDataOrder[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrNotFound
}

// UpdateOrderItemById
func (m *OrderRepositoryMock) UpdateOrderItemById(ctx context.Context, client *ent.Client, orderItemId string, payload *dto.OrderItem) (*ent.OrderItem, error) {
	m.Lock()

	_, err := uuid.Parse(orderItemId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	for key, data := range m.mockDataOrderItem {
		if key == orderItemId {
			u := data
			u.PurchasedName = *payload.PurchasedName
			u.PurchasedPrice = *payload.PurchasedPrice
			u.Quantity = *payload.Quantity

			m.mockDataOrderItem[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrNotFound
}

// SoftDeleteOrderById
func (m *OrderRepositoryMock) SoftDeleteOrderById(ctx context.Context, client *ent.Client, orderId string) (*ent.Order, error) {
	m.Lock()

	_, err := uuid.Parse(orderId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	for key, data := range m.mockDataOrder {
		if key == orderId {
			u := data
			if u.IsArchived {
				return nil, constants.ErrBadRequest
			}
			u.IsArchived = true

			m.mockDataOrder[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrNotFound
}

// DeleteOrderItemById
func (m *OrderRepositoryMock) DeleteOrderItemById(ctx context.Context, client *ent.Client, orderItemId string) (bool, error) {
	m.Lock()

	_, err := uuid.Parse(orderItemId)
	if err != nil {
		return false, constants.ErrBadRequest
	}

	delete(m.mockDataOrderItem, orderItemId)
	return true, nil
}
