package repository

import (
	"context"
	"net/mail"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/ent/user"
	"sthl/storage"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IUserRepository interface {
	WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error
	CreateUser(ctx context.Context, client *ent.Client, payload *dto.CreateUserMappedDto) (*ent.User, error)
	GetUserById(ctx context.Context, client *ent.Client, userId string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error)
	UpdateUserPasswordById(ctx context.Context, client *ent.Client, userId string, hashedPw string) (*ent.User, error)
	// for test
	GetUsers(ctx context.Context, client *ent.Client, payload *dto.QueryUsersDto) (*dto.QueryUsersResponseDto, error)
}

type UserRepository struct {
	logger *zap.Logger
}

func NewUserRepository(logger *zap.Logger) IUserRepository {
	return &UserRepository{
		logger: logger,
	}
}

func (userRepo *UserRepository) WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	return storage.WithTx(ctx, userRepo.logger, client, fn)
}

// CreateUser
func (userRepo *UserRepository) CreateUser(ctx context.Context, client *ent.Client, payload *dto.CreateUserMappedDto) (*ent.User, error) {
	result, err := client.User.Create().
		SetEmail(*payload.Email).
		SetHashedPw(*payload.HashedPw).
		Save(ctx)
	if err != nil {
		userRepo.logger.Info("fail to client.User.Create", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// GetUserById
func (userRepo *UserRepository) GetUserById(ctx context.Context, client *ent.Client, userId string) (*ent.User, error) {
	// parse userId to uuid
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		userRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call ent client to get
	result, err := client.User.Get(ctx, userUuid)
	if err != nil {
		userRepo.logger.Info("fail to client.User.Get", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// GetUserByEmail
func (userRepo *UserRepository) GetUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	// validate email
	_, err := mail.ParseAddress(email)
	if err != nil {
		userRepo.logger.Info("fail to parseAddress", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call ent client to query
	result, err := client.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		userRepo.logger.Info("fail to client.User.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// UpdateUserPasswordById
func (userRepo *UserRepository) UpdateUserPasswordById(ctx context.Context, client *ent.Client, userId string, hashedPw string) (*ent.User, error) {
	// parse userId to uuid
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		userRepo.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call ent client to updateOneID
	result, err := client.User.UpdateOneID(userUuid).
		SetHashedPw(hashedPw).
		Save(ctx)
	if err != nil {
		userRepo.logger.Info("fail to client.User.UpdateOneID", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}
	return result, nil
}

// GetUsers
func (userRepo *UserRepository) GetUsers(ctx context.Context, client *ent.Client, payload *dto.QueryUsersDto) (*dto.QueryUsersResponseDto, error) {
	page := payload.Page
	limit := payload.Limit
	offset := (page - 1) * limit

	total, err := client.User.Query().Count(ctx)
	if err != nil {
		userRepo.logger.Info("fail to count total", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	// call ent client to updateOneID
	result, err := client.User.Query().
		Order(ent.Desc(user.FieldCreatedAt)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		userRepo.logger.Info("fail to client.User.Query", zap.Error(err))
		return nil, handleEntRepoErr(err)
	}

	pagingResp := dto.NewPagingResponse(page, limit, total)
	data := dto.NewQueryUsersResponseDto(result, *pagingResp)
	return data, nil
}
