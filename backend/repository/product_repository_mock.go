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

type ProductRepositoryMock struct {
	mockData map[string]ent.Product
	mu       sync.Mutex
}

func NewProductRepositoryMock() IProductRepository {
	return &ProductRepositoryMock{
		mockData: map[string]ent.Product{},
	}
}

func newProductSchemaMock(userId uuid.UUID, payload *dto.CreateProductDtoMappedDto) *ent.Product {
	t := time.Now()
	return &ent.Product{
		ID:          uuid.New(),
		CreatedAt:   t,
		UpdatedAt:   t,
		UserID:      userId,
		Name:        *payload.Name,
		Price:       *payload.Price,
		Quantity:    *payload.Quantity,
		Description: *payload.Description,
		Status:      *payload.Status,
		IsArchived:  false,
	}
}

func (m *ProductRepositoryMock) Lock() {
	m.mu.Lock()
	defer m.mu.Unlock()
}
func (m *ProductRepositoryMock) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTxTest(ctx, nil, client, fn)
}

// ****

// CreateProduct
func (m *ProductRepositoryMock) CreateProduct(
	ctx context.Context, client *ent.Client, userId string, payload *dto.CreateProductDtoMappedDto) (*ent.Product, error) {
	m.Lock()
	for _, data := range m.mockData {
		if data.Name == *payload.Name {
			return nil, constants.ErrExisted
		}
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, constants.ErrBadRequest
	}
	result := newProductSchemaMock(userUuid, payload)
	m.mockData[result.ID.String()] = *result
	return result, nil
}

// GetProductsTotalByUserId
func (m *ProductRepositoryMock) GetProductsTotalByUserId(ctx context.Context, client *ent.Client, userId string) (int, error) {
	var count int = 0
	for _, u := range m.mockData {
		if u.UserID.String() == userId {
			count++
		}
	}
	return count, nil
}

// GetProducts
func (m *ProductRepositoryMock) GetProducts(
	ctx context.Context, client *ent.Client, userId string, payload *dto.QueryProductsDto) (*dto.QueryProductsResponseDto, error) {
	m.Lock()
	var productSlice, subSlice []*ent.Product
	for _, u := range m.mockData {
		if u.UserID.String() == userId {
			productSlice = append(productSlice, &u)
		}
	}
	sort.Slice(productSlice, func(i, j int) bool {
		return productSlice[i].CreatedAt.Before(productSlice[j].CreatedAt)
	})

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit
	subSlice = productSlice[offset:]
	if len(productSlice) > offset+limit-1 {
		subSlice = productSlice[offset : offset+limit-1]
	}

	pagingResp := dto.NewPagingResponse(payload.Page, payload.Limit, len(productSlice))
	result := dto.NewQueryProductsResponseDto(subSlice, *pagingResp)
	return result, nil
}

// GetProductById
func (m *ProductRepositoryMock) GetProductById(
	ctx context.Context, client *ent.Client, productId string) (*ent.Product, error) {
	m.Lock()

	value := m.mockData[productId]
	if utils.IsEmpty(value) {
		return nil, constants.ErrNotFound
	}
	return &value, nil
}

// UpdateProductById
func (m *ProductRepositoryMock) UpdateProductById(
	ctx context.Context, client *ent.Client, productId string, payload *dto.UpdateProductDto) (*ent.Product, error) {
	m.Lock()
	for key, data := range m.mockData {
		if key == productId {
			u := data

			u.Name = *payload.Name
			u.Price = *payload.Price
			u.Quantity = *payload.Quantity
			u.Description = *payload.Description
			u.Status = *payload.Status

			m.mockData[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrNotFound
}

// SoftDeleteProductById
func (m *ProductRepositoryMock) SoftDeleteProductById(
	ctx context.Context, client *ent.Client, productId string) (*ent.Product, error) {
	m.Lock()

	for key, data := range m.mockData {
		if key == productId {
			u := data
			if u.IsArchived {
				return nil, constants.ErrBadRequest
			}
			u.IsArchived = true

			m.mockData[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrBadRequest
}
