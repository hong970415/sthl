package service

import (
	"context"
	"sthl/authentication"
	"sthl/constants"
	"sthl/dto"
	"sthl/ent"
	"sthl/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IUserService interface {
	// public
	Signup(ctx context.Context, payload *dto.CreateUserDto) (*ent.User, error)
	Login(ctx context.Context, payload *dto.LoginDto) (*authentication.Passport, error)
	// private
	RefreshAccessToken(ctx context.Context, userId string, payload *dto.RefreshAccessTokenDto) (*authentication.Passport, error)
	GetUserById(ctx context.Context, userId string) (*ent.User, error)
	UpdateUserPasswordById(ctx context.Context, userId string, payload *dto.UpdateUserPasswordDto) (*ent.User, error)
	// internal
	authentication.Authenticator
	// testing
	GetUsers(ctx context.Context, payload *dto.QueryUsersDto) (*dto.QueryUsersResponseDto, error)
}
type UserService struct {
	logger   *zap.Logger
	client   *ent.Client
	userRepo repository.IUserRepository
}

func NewUserService(logger *zap.Logger, client *ent.Client, userRepo repository.IUserRepository) IUserService {
	return &UserService{
		logger:   logger,
		client:   client,
		userRepo: userRepo,
	}
}

// Signup
func (userSvc *UserService) Signup(ctx context.Context, payload *dto.CreateUserDto) (*ent.User, error) {
	// validate
	err := payload.Validate()
	if err != nil {
		userSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// hash password
	hashPw, err := authentication.HashPassword(*payload.Password)
	if err != nil {
		userSvc.logger.Info("fail to hashPassword", zap.Error(err))
		return nil, constants.ErrInternalServer
	}

	// mapping
	hashedPayload := payload.MapToSchema(hashPw)

	// call repo to createUser with tx
	var result *ent.User
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()
		rs, err := userSvc.userRepo.CreateUser(ctx, txc, hashedPayload)
		if err != nil {
			return err
		}
		result = rs
		return nil
	}
	err = userSvc.userRepo.WithTx(ctx, userSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Login
func (userSvc *UserService) Login(ctx context.Context, payload *dto.LoginDto) (*authentication.Passport, error) {
	// validate
	err := payload.Validate()
	if err != nil {
		userSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to createUser
	result, err := userSvc.userRepo.GetUserByEmail(ctx, userSvc.client, *payload.Email)
	if err != nil {
		return nil, err
	}

	// compare password
	err = authentication.CompareHashPassword(result.HashedPw, *payload.Password)
	if err != nil {
		return nil, constants.ErrBadRequest
	}

	// generate pp
	pp, err := authentication.GeneratePassport(
		constants.AccessTokenDuration, constants.RefreshTokenDuration, result.ID.String())
	if err != nil {
		return nil, constants.ErrInternalServer
	}
	return pp, nil
}

// RefreshAccessToken
func (userSvc *UserService) RefreshAccessToken(
	ctx context.Context, userId string, payload *dto.RefreshAccessTokenDto) (*authentication.Passport, error) {
	// validate
	err := payload.Validate()
	if err != nil {
		userSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// verify refresh token
	claim, err := authentication.VerifyJwtToken(*payload.RefreshToken)
	if err != nil {
		userSvc.logger.Info("fail to verifyJwtToken", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// check at and rt equal userId
	if userId != claim.UserId {
		userSvc.logger.Info("userId not equal claim.UserId")
		return nil, constants.ErrBadRequest
	}

	// pass, re-generate pp
	pp, err := authentication.GeneratePassport(constants.AccessTokenDuration, constants.RefreshTokenDuration, userId)
	if err != nil {
		userSvc.logger.Info("fail to generatePassport", zap.Error(err))
		return nil, constants.ErrInternalServer
	}
	return pp, nil
}

// GetUserById
func (userSvc *UserService) GetUserById(ctx context.Context, userId string) (*ent.User, error) {
	// parse userId to uuid
	_, err := uuid.Parse(userId)
	if err != nil {
		userSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// check is zero uuid
	if userId == uuid.Nil.String() {
		userSvc.logger.Info("userId is zero uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// call repo to get user
	result, err := userSvc.userRepo.GetUserById(ctx, userSvc.client, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUserPasswordById
func (userSvc *UserService) UpdateUserPasswordById(
	ctx context.Context, userId string, payload *dto.UpdateUserPasswordDto) (*ent.User, error) {
	// parse userId to uuid
	_, err := uuid.Parse(userId)
	if err != nil {
		userSvc.logger.Info("fail to parse userId to uuid", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// validate
	err = payload.Validate()
	if err != nil {
		userSvc.logger.Info("fail to validate", zap.Error(err))
		return nil, constants.ErrBadRequest
	}

	// update with transaction
	var result *ent.User
	txFunc := func(tx *ent.Tx) error {
		// extract tx client as ent.cient
		txc := tx.Client()

		// call repo to getUserById
		found, err := userSvc.userRepo.GetUserById(ctx, txc, userId)
		if err != nil {
			return err
		}

		// compare password
		err = authentication.CompareHashPassword(found.HashedPw, *payload.CurrentPassword)
		if err != nil {
			userSvc.logger.Info("fail to validate", zap.Error(err))
			return constants.ErrBadRequest
		}

		// hash new pw
		hashedNewPw, err := authentication.HashPassword(*payload.NewPassword)
		if err != nil {
			userSvc.logger.Info("fail to new hashPassword", zap.Error(err))
			return constants.ErrBadRequest
		}

		// call repo to updateUserPasswordById
		updateResult, err := userSvc.userRepo.UpdateUserPasswordById(ctx, txc, userId, hashedNewPw)
		result = updateResult
		if err != nil {
			return constants.ErrBadRequest
		}
		return nil
	}
	err = userSvc.userRepo.WithTx(ctx, userSvc.client, txFunc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Authenticate
func (userSvc *UserService) Authenticate(ctx context.Context, token string) (string, error) {
	result, err := authentication.VerifyJwtToken(token)
	if err != nil {
		userSvc.logger.Info("fail to verifyJwtToken", zap.Error(err))
		return "", constants.ErrBadRequest
	}
	return result.UserId, nil
}

// GetUsers
func (userSvc *UserService) GetUsers(ctx context.Context, payload *dto.QueryUsersDto) (*dto.QueryUsersResponseDto, error) {
	ensuredPaging := payload.Ensure()
	ensuredPayload := dto.NewQueryUsersDto(*ensuredPaging)

	result, err := userSvc.userRepo.GetUsers(ctx, userSvc.client, ensuredPayload)
	if err != nil {
		return nil, err
	}
	return result, nil
}
