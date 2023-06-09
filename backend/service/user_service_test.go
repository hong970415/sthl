package service

import (
	"context"
	"sthl/authentication"
	"sthl/dto"
	"sthl/ent"
	"sthl/logger"
	"sthl/repository"
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// userServiceTestSetup
func userServiceTestSetup(ctx context.Context, t *testing.T) (*assert.Assertions, IUserService) {
	t.Setenv("JWT_SECRET", "test_value")
	assert := assert.New(t)
	// dependency init
	// zapLogger, err := logger.NewDevInfoZapLogger()
	zapLogger, err := logger.NewDevErrorZapLogger()
	assert.NotEmpty(zapLogger)
	assert.NoError(err)
	userRepo := repository.NewUserRepositoryMock()
	userService := NewUserService(zapLogger, nil, userRepo)
	assert.NotEmpty(userRepo)
	assert.NotEmpty(userService)
	return assert, userService
}

// ****Test_Signup
type signupTestCase struct {
	name  string
	input *dto.CreateUserDto
	exec  func(*ent.User, error)
}

func Test_Signup(t *testing.T) {
	ctx := context.TODO()
	assert, userService := userServiceTestSetup(ctx, t)

	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	duplicateEmail := gofakeit.Email()
	assert.NotEmpty(validCreateUserDto)

	testCases := []signupTestCase{
		{
			name:  "signup with valid param",
			input: validCreateUserDto,
			exec: func(result *ent.User, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name: "signup with invalid param, username already existed",
			input: dto.NewCreateUserDto(
				utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name: "signup with invalid param, username and password are same",
			input: dto.NewCreateUserDto(
				utils.PtrOf(duplicateEmail), utils.PtrOf(duplicateEmail)),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name: "signup with invalid param, not email",
			input: dto.NewCreateUserDto(
				utils.PtrOf(gofakeit.Name()), utils.PtrOf(gofakeit.Name())),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userService.Signup(ctx, test.input))
		})
	}
}

// ****Test_Login
type loginTestCase struct {
	name  string
	input *dto.LoginDto
	exec  func(*authentication.Passport, error)
}

func Test_Login(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// pre signup user
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser, err := userSvc.Signup(ctx, validCreateUserDto)
	assert.NotEmpty(validUser)
	assert.NoError(err)

	testCases := []loginTestCase{
		{
			name: "login with valid info",
			input: dto.NewLoginDto(
				utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf(*validCreateUserDto.Password)),
			exec: func(result *authentication.Passport, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name: "login with non existed user",
			//todo gofakeit
			input: dto.NewLoginDto(
				utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name: "login with wrong password",
			input: dto.NewLoginDto(
				utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf("wrongpassword")),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name: "login with same username and password",
			input: dto.NewLoginDto(
				utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf(*validCreateUserDto.Email)),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name: "login with invalid email",
			input: dto.NewLoginDto(
				utils.PtrOf(gofakeit.Name()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.Login(ctx, test.input))
		})
	}
}

// ****Test_RefreshAccessToken
type refreshAccessTokenCase struct {
	name   string
	userId string
	input  *dto.RefreshAccessTokenDto
	exec   func(*authentication.Passport, error)
}

func Test_RefreshAccessToken(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// pre signup a user and login user
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser, err := userSvc.Signup(ctx, validCreateUserDto)
	assert.NotEmpty(validUser)
	assert.NoError(err)
	validPp, err := userSvc.Login(ctx,
		dto.NewLoginDto(utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf(*validCreateUserDto.Password)))
	assert.NotEmpty(validPp)
	assert.NoError(err)

	testCases := []refreshAccessTokenCase{
		{
			name:   "with valid token",
			userId: validUser.ID.String(),
			input:  dto.NewRefreshAccessTokenDto(&validPp.RefreshToken),
			exec: func(result *authentication.Passport, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:   "with invalid token",
			userId: validUser.ID.String(),
			input:  dto.NewRefreshAccessTokenDto(utils.PtrOf("qwdwefwef")),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "with invalid userId",
			userId: "invailduserid",
			input:  dto.NewRefreshAccessTokenDto(&validPp.RefreshToken),
			exec: func(result *authentication.Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.RefreshAccessToken(ctx, test.userId, test.input))
		})
	}
}

// ****Test_GetUserById
type getUserByIdTestCase struct {
	name  string
	input string
	exec  func(*ent.User, error)
}

func Test_GetUserById(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// pre signup user
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser, err := userSvc.Signup(ctx, validCreateUserDto)
	assert.NotEmpty(validUser)
	assert.NoError(err)

	testCases := []getUserByIdTestCase{
		{
			name:  "get with valid userId",
			input: validUser.ID.String(),
			exec: func(result *ent.User, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "get with invalid userId",
			input: "123",
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:  "get with invalid zero userId, 00000000-0000-0000-0000-000000000000",
			input: uuid.Nil.String(),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.GetUserById(ctx, test.input))
		})
	}
}

// ****Test_UpdateUserPasswordById
type updateUserPasswordByIdTestCase struct {
	name   string
	userId string
	input  *dto.UpdateUserPasswordDto
	exec   func(*ent.User, error)
}

func Test_UpdateUserPasswordById(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// pre signup user
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser, err := userSvc.Signup(ctx, validCreateUserDto)
	assert.NotEmpty(validUser)
	assert.NoError(err)

	testCases := []updateUserPasswordByIdTestCase{
		{
			name:   "update with invalid userId",
			userId: "123",
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(*validCreateUserDto.Password),
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "update with invalid zero userId",
			userId: uuid.Nil.String(),
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(*validCreateUserDto.Password),
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "update with same old and new pw",
			userId: validUser.ID.String(),
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(*validCreateUserDto.Password),
				utils.PtrOf(*validCreateUserDto.Password)),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "update with wrong pw",
			userId: validUser.ID.String(),
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)),
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *ent.User, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:   "update with correct pw, sccuess update",
			userId: validUser.ID.String(),
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(*validCreateUserDto.Password),
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(result *ent.User, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.UpdateUserPasswordById(ctx, test.userId, test.input))
		})
	}
}

// ****Test_Authenticate
type authenticateTestCase struct {
	name  string
	input string
	exec  func(string, error)
}

func Test_Authenticate(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// pre signup a user and login user
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser, err := userSvc.Signup(ctx, validCreateUserDto)
	assert.NotEmpty(validUser)
	assert.NoError(err)
	validPp, err := userSvc.Login(ctx,
		dto.NewLoginDto(utils.PtrOf(*validCreateUserDto.Email), utils.PtrOf(*validCreateUserDto.Password)))
	assert.NotEmpty(validPp)
	assert.NoError(err)

	testCases := []authenticateTestCase{
		{
			name:  "valid tokenString",
			input: validPp.AccessToken,
			exec: func(result string, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "empty tokenString",
			input: "",
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:  "invalid tokenString",
			input: "qwdqwd",
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.Authenticate(ctx, test.input))
		})
	}
}

// ****Test_GetUsers, for testing
type getUsersTestCase struct {
	name  string
	input *dto.QueryUsersDto
	exec  func(*dto.QueryUsersResponseDto, error)
}

func Test_GetUsers(t *testing.T) {
	ctx := context.TODO()
	assert, userSvc := userServiceTestSetup(ctx, t)

	// // pre signup 2 user
	validCreateUser1Dto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validCreateUser2Dto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	validUser1, err := userSvc.Signup(ctx, validCreateUser1Dto)
	assert.NotEmpty(validUser1)
	assert.NoError(err)
	validUser2, err := userSvc.Signup(ctx, validCreateUser2Dto)
	assert.NotEmpty(validUser2)
	assert.NoError(err)

	testCases := []getUsersTestCase{
		{
			name:  "invalid paging",
			input: dto.NewQueryUsersDto(*dto.NewPaging(-1, 2, "")),
			exec: func(result *dto.QueryUsersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "invalid paging, limit exceed 1000",
			input: dto.NewQueryUsersDto(*dto.NewPaging(1, 1001, "")),
			exec: func(result *dto.QueryUsersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "valid paging",
			input: dto.NewQueryUsersDto(*dto.NewPaging(1, 10, "")),
			exec: func(result *dto.QueryUsersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "valid paging, limit",
			input: dto.NewQueryUsersDto(*dto.NewPaging(1, 1, "")),
			exec: func(result *dto.QueryUsersResponseDto, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(userSvc.GetUsers(ctx, test.input))
		})
	}
}
