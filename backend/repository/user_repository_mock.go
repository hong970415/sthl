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

type UserRepositoryMock struct {
	mockData map[string]ent.User
	mu       sync.Mutex
}

func NewUserRepositoryMock() IUserRepository {
	return &UserRepositoryMock{
		mockData: map[string]ent.User{},
	}
}

func newUserSchemaMock(payload *dto.CreateUserMappedDto) *ent.User {
	t := time.Now()
	return &ent.User{
		ID:            uuid.New(),
		CreatedAt:     t,
		UpdatedAt:     t,
		Email:         *payload.Email,
		HashedPw:      *payload.HashedPw,
		EmailVerified: false,
	}
}

func (m *UserRepositoryMock) Lock() {
	m.mu.Lock()
	defer m.mu.Unlock()
}

func (m *UserRepositoryMock) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTxTest(ctx, nil, client, fn)
}

// ****

// CreateUser
func (m *UserRepositoryMock) CreateUser(ctx context.Context, client *ent.Client, payload *dto.CreateUserMappedDto) (*ent.User, error) {
	m.Lock()
	for _, data := range m.mockData {
		if data.Email == *payload.Email {
			return nil, constants.ErrExisted
		}
	}
	result := newUserSchemaMock(payload)
	m.mockData[result.ID.String()] = *result
	return result, nil
}

// GetUserById
func (m *UserRepositoryMock) GetUserById(ctx context.Context, client *ent.Client, userId string) (*ent.User, error) {
	m.Lock()
	value := m.mockData[userId]
	if utils.IsEmpty(value) {
		return nil, constants.ErrNotFound
	}
	return &value, nil
}

// GetUserByEmail
func (m *UserRepositoryMock) GetUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	m.Lock()
	for _, data := range m.mockData {
		if data.Email == email {
			return &data, nil
		}
	}
	return nil, constants.ErrNotFound
}

// UpdateUserPasswordById
func (m *UserRepositoryMock) UpdateUserPasswordById(ctx context.Context, client *ent.Client, userId string, hashedPw string) (*ent.User, error) {
	m.Lock()
	for key, data := range m.mockData {
		if key == userId {
			u := data
			u.HashedPw = hashedPw
			m.mockData[key] = u
			return &u, nil
		}
	}
	return nil, constants.ErrNotFound
}

// for test

// GetUsers
func (m *UserRepositoryMock) GetUsers(ctx context.Context, client *ent.Client, payload *dto.QueryUsersDto) (*dto.QueryUsersResponseDto, error) {
	m.Lock()
	var userSlice, subSlice []*ent.User
	for _, u := range m.mockData {
		userSlice = append(userSlice, &u)
	}
	sort.Slice(userSlice, func(i, j int) bool {
		return userSlice[i].CreatedAt.Before(userSlice[j].CreatedAt)
	})

	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit
	subSlice = userSlice[offset:]
	if len(userSlice) > offset+limit-1 {
		subSlice = userSlice[offset : offset+limit-1]
	}

	pagingResp := dto.NewPagingResponse(payload.Page, payload.Limit, len(userSlice))
	result := dto.NewQueryUsersResponseDto(subSlice, *pagingResp)
	return result, nil
}
