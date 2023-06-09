package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sthl/authentication"
	"sthl/config"
	"sthl/dto"
	"sthl/ent"
	"sthl/logger"
	"sthl/repository"
	"sthl/service"
	"sthl/storage"
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// ****helpers****

// generateHttpTestRequestBody
func generateHttpTestRequestBody(assert *assert.Assertions, input any) io.Reader {
	jsonBytes, err := json.Marshal(input)
	reqBody := bytes.NewReader(jsonBytes)
	assert.NoError(err)
	assert.NotEmpty(reqBody)
	return reqBody
}

// executeHttpTestRequest
func executeHttpTestRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	req.Header.Set("X-Request-Id", "Test-Header")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

// ****setup****

// setTestEnv
func setTestEnv(t *testing.T) {
	t.Setenv("NODE_ENV", "develop")
	t.Setenv("PORT", "4000")
	t.Setenv("DB_DOMAIN", "127.0.0.1")
	t.Setenv("DB_USER", "postgres")
	t.Setenv("DB_PASSWORD", "postgres")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("JWT_SECRET", "testsecret")
	t.Setenv("ALLOW_ORIGIN", "http://localhost:3000")
	t.Setenv("AWS_ACCESS_KEY_ID", "test")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test")
	t.Setenv("AWS_REGION", "ap-east-1")
	t.Setenv("S3_PATH", "http://localhost:4566")
	t.Setenv("VERSION", "v1")
}

// handlersTestSetup
func handlersTestSetup(ctx context.Context, t *testing.T) (*assert.Assertions, *chi.Mux) {
	assert := assert.New(t)
	// zapLogger, err := logger.NewDevInfoZapLogger()
	zapLogger, err := logger.NewDevErrorZapLogger()
	assert.NotEmpty(zapLogger)
	assert.NoError(err)

	// repos
	var userRepo repository.IUserRepository
	var productRepo repository.IProductRepository
	var orderRepo repository.IOrderRepository
	var siteuiRepo repository.ISiteUiRepository
	var imageInfoRepo repository.IImgInfoRepository

	// services
	var userSvc service.IUserService
	var productSvc service.IProductService
	var orderSvc service.IOrderService
	var siteuiSvc service.ISiteUiService
	var albumSvc service.IAlbumService

	setTestEnv(t)
	cfg, ok := config.NewConfig(zapLogger)
	assert.NotEmpty(cfg)
	assert.Equal(true, ok)

	if testing.Short() {
		// case unit test
		userRepo = repository.NewUserRepositoryMock()
		productRepo = repository.NewProductRepositoryMock()
		orderRepo = repository.NewOrderRepositoryMock()
		siteuiRepo = nil
		imageInfoRepo = nil

		userSvc = service.NewUserService(zapLogger, nil, userRepo)
		productSvc = service.NewProductService(zapLogger, nil, userRepo, productRepo)
		orderSvc = service.NewOrderService(zapLogger, nil, userRepo, productRepo, orderRepo)
		siteuiSvc = service.NewSiteUiService(zapLogger, nil, userRepo, siteuiRepo)
		albumSvc = service.NewAlbumService(zapLogger, nil, nil, imageInfoRepo)
	} else {
		// case integration test

		// config -> new postgres test container -> remove container after test
		rs, err := storage.NewPostgresTestContainer(zapLogger, cfg)
		assert.NotEmpty(rs)
		assert.NoError(err)
		t.Cleanup(func() {
			err := rs.Terminate(ctx)
			assert.NoError(err)
		})

		// connect to postgres test container
		dbclient, err := storage.NewPostgresDb(zapLogger, cfg)
		assert.NotEmpty(dbclient)
		assert.NoError(err)

		userRepo = repository.NewUserRepository(zapLogger)
		productRepo = repository.NewProductRepository(zapLogger)
		orderRepo = repository.NewOrderRepository(zapLogger)
		siteuiRepo = nil
		imageInfoRepo = nil

		userSvc = service.NewUserService(zapLogger, dbclient, userRepo)
		productSvc = service.NewProductService(zapLogger, dbclient, userRepo, productRepo)
		orderSvc = service.NewOrderService(zapLogger, dbclient, userRepo, productRepo, orderRepo)
		siteuiSvc = service.NewSiteUiService(zapLogger, dbclient, userRepo, siteuiRepo)
		albumSvc = service.NewAlbumService(zapLogger, dbclient, nil, imageInfoRepo)
	}

	hdlers := NewHandler(zapLogger, userSvc, productSvc, orderSvc, siteuiSvc, albumSvc)
	r := NewChiRouter(zapLogger, cfg, userSvc, hdlers)
	return assert, r
}

// pre signup user
func preSignupUser(assert *assert.Assertions, r *chi.Mux) (*dto.CreateUserDto, *ent.User) {
	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	b := generateHttpTestRequestBody(assert, *validCreateUserDto)
	req, err := http.NewRequest("POST", "/api/v1/users", b)
	assert.NotEmpty(req)
	assert.NoError(err)

	preuserRr := executeHttpTestRequest(req, r)
	var rs utils.ResponseMessage[ent.User]
	err = json.Unmarshal(preuserRr.Body.Bytes(), &rs)
	assert.NotEmpty(rs)
	assert.NoError(err)

	assert.Equal(http.StatusCreated, preuserRr.Code)

	return validCreateUserDto, rs.Data
}

// pre login to get pp
func preLoginUser(assert *assert.Assertions, r *chi.Mux, payload *dto.LoginDto) *authentication.Passport {
	assert.NotEmpty(payload)
	loginBody := generateHttpTestRequestBody(assert, *payload)
	loginReq, err := http.NewRequest("POST", "/api/v1/users/login", loginBody)
	assert.NotEmpty(loginReq)
	assert.NoError(err)
	loginguserRr := executeHttpTestRequest(loginReq, r)
	var logingResp utils.ResponseMessage[authentication.Passport]
	err = json.Unmarshal(loginguserRr.Body.Bytes(), &logingResp)
	assert.NotEmpty(logingResp)
	assert.NoError(err)
	assert.Equal(http.StatusOK, loginguserRr.Code)
	return logingResp.Data
}

// preSignupLoginUser
func preSignupLoginUser(assert *assert.Assertions, r *chi.Mux) (*authentication.Passport, *dto.LoginDto, *ent.User) {
	validCreateUser, validUser := preSignupUser(assert, r)
	validLoginUser := dto.NewLoginDto(validCreateUser.Email, validCreateUser.Password)
	validPp := preLoginUser(assert, r, validLoginUser)
	return validPp, validLoginUser, validUser
}

// ****tests****

// Test_HandlePing
func Test_HandlePing(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	req, err := http.NewRequest("GET", "/api/v1/ping", nil)
	assert.NotEmpty(req)
	assert.NoError(err)

	rr := executeHttpTestRequest(req, r)
	assert.Equal(http.StatusOK, rr.Code)
}

// ****User

// Test_HandleSignup
type handleSignupTestCase struct {
	name  string
	input *dto.CreateUserDto
	exec  func(*httptest.ResponseRecorder)
}

func Test_HandleSignup(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validCreateUserDto := dto.NewCreateUserDto(
		utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))

	testCases := []handleSignupTestCase{
		{
			name:  "signup with valid param",
			input: validCreateUserDto,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusCreated, rr.Code)
			},
		},
		{
			name:  "signup with invalid param, existed email",
			input: validCreateUserDto,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusBadRequest, rr.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := generateHttpTestRequestBody(assert, *test.input)
			req, err := http.NewRequest("POST", "/api/v1/users", b)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_Login
type loginTestCase struct {
	name  string
	input *dto.LoginDto
	exec  func(*httptest.ResponseRecorder)
}

func Test_HandleLogin(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validCreateUserDto, _ := preSignupUser(assert, r)

	testCases := []loginTestCase{
		{
			name:  "login with valid param",
			input: dto.NewLoginDto(validCreateUserDto.Email, validCreateUserDto.Password),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[authentication.Passport]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
			},
		},
		{
			name:  "login with invalid param, email",
			input: dto.NewLoginDto(utils.PtrOf(gofakeit.Email()), validCreateUserDto.Password),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[authentication.Passport]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusNotFound, rr.Code)
			},
		},
		{
			name:  "login with invalid param, password",
			input: dto.NewLoginDto(validCreateUserDto.Email, utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6))),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[authentication.Passport]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusBadRequest, rr.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := generateHttpTestRequestBody(assert, *test.input)
			req, err := http.NewRequest("POST", "/api/v1/users/login", b)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleRefreshAccessToken
type handleRefreshAccessTokenTestCase struct {
	name        string
	headerToken string
	input       *dto.RefreshAccessTokenDto
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleRefreshAccessToken(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, _, _ := preSignupLoginUser(assert, r)

	testCases := []handleRefreshAccessTokenTestCase{
		{
			name:        "login with valid param",
			headerToken: validPp.AccessToken,
			input:       dto.NewRefreshAccessTokenDto(&validPp.RefreshToken),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[authentication.Passport]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
			},
		},
		{
			name:        "login with invalid param, header token",
			headerToken: "qwdqwdwdqf",
			input:       dto.NewRefreshAccessTokenDto(&validPp.RefreshToken),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[any]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
		{
			name:        "login with invalid param, refresh token",
			headerToken: validPp.AccessToken,
			input:       dto.NewRefreshAccessTokenDto(utils.PtrOf("invalidrt")),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[any]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusBadRequest, rr.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := generateHttpTestRequestBody(assert, *test.input)
			req, err := http.NewRequest("POST", "/api/v1/users/refreshToken", b)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleGetMe
type handleGetMeTestCase struct {
	name        string
	headerToken string
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleGetMe(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, _, _ := preSignupLoginUser(assert, r)

	testCases := []handleGetMeTestCase{
		{
			name:        "getme with valid param",
			headerToken: validPp.AccessToken,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
			},
		},
		{
			name:        "getme with invalid param, access token",
			headerToken: "invalidtoken",
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/api/v1/users/me", nil)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleUpdateUserPasswordById
type handleUpdateUserPasswordByIdTestCase struct {
	name        string
	headerToken string
	input       *dto.UpdateUserPasswordDto
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleUpdateUserPasswordById(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, validCreateUser, _ := preSignupLoginUser(assert, r)
	newPw := gofakeit.Password(true, true, true, true, false, 6)
	testCases := []handleUpdateUserPasswordByIdTestCase{
		{
			name:        "update with valid param",
			headerToken: validPp.AccessToken,
			input:       dto.NewUpdateUserPasswordDto(validCreateUser.Password, &newPw),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)

				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
			},
		},
		{
			name:        "update with invalid param, headerToken",
			headerToken: "invalidToken",
			input:       dto.NewUpdateUserPasswordDto(validCreateUser.Password, &newPw),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
		{
			name:        "update with invalid param, headerToken",
			headerToken: validPp.AccessToken,
			input: dto.NewUpdateUserPasswordDto(
				utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)), &newPw),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusBadRequest, rr.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := generateHttpTestRequestBody(assert, test.input)
			req, err := http.NewRequest("PUT", "/api/v1/users/me/pw", b)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleGetUsers
type handleGetUsersTestCase struct {
	name        string
	headerToken string
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleGetUsers(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, _, _ := preSignupLoginUser(assert, r)

	testCases := []handleGetUsersTestCase{
		{
			name:        "get with valid param",
			headerToken: validPp.AccessToken,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[dto.QueryUsersResponseDto]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
				assert.Equal(1, rs.Data.Total)
			},
		},
		{
			name:        "get with invalid param, headertoken",
			headerToken: "invalid",
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[dto.QueryUsersResponseDto]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/api/v1/users", nil)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleGetUserById
type handleGetUserByIdTestCase struct {
	name        string
	headerToken string
	input       string
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleGetUserById(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, _, validUser := preSignupLoginUser(assert, r)
	// _, _, validUser2 := preSignupLoginUser(assert, r)

	testCases := []handleGetUserByIdTestCase{
		{
			name:        "get with valid param",
			headerToken: validPp.AccessToken,
			input:       validUser.ID.String(),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)

				assert.NotEmpty(rs.Data)
			},
		},
		{
			name:        "get with invalid param, headerToken",
			headerToken: "wqdqwd",
			input:       validUser.ID.String(),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.User]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
		// {
		// 	name:        "get with invalid param, headerToken",
		// 	headerToken: validPp.AccessToken,
		// 	input:       validUser2.ID.String(),
		// 	exec: func(rr *httptest.ResponseRecorder) {
		// 		var rs utils.ResponseMessage[ent.User]
		// 		err := json.Unmarshal(rr.Body.Bytes(), &rs)
		// 		assert.NotEmpty(rs)
		// 		assert.NoError(err)
		// 		assert.Equal(http.StatusBadRequest, rr.Code)
		// 	},
		// },
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v1/users/%s", test.input)
			req, err := http.NewRequest("GET", path, nil)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// ****Product

// Test_HandleCreateProduct
type handleCreateProductTestCase struct {
	name        string
	headerToken string
	input       *dto.CreateProductDto
	exec        func(*httptest.ResponseRecorder)
}

func Test_HandleCreateProduct(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	validPp, _, _ := preSignupLoginUser(assert, r)
	validCreateProductDto := dto.NewCreateProductDto(
		utils.PtrOf(gofakeit.Fruit()),
		utils.PtrOf(gofakeit.Float64Range(1.0, 1000000.0)),
		utils.PtrOf(int32(gofakeit.IntRange(0, 1000000))),
		utils.PtrOf(gofakeit.LetterN(100)),
		utils.PtrOf(gofakeit.LetterN(100)))

	testCases := []handleCreateProductTestCase{
		{
			name:        "create with valid param",
			headerToken: validPp.AccessToken,
			input:       validCreateProductDto,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.Product]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusCreated, rr.Code)
			},
		},
		{
			name:        "create with invalid headertoken",
			headerToken: validPp.AccessToken + "wrong",
			input:       validCreateProductDto,
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[ent.Product]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusUnauthorized, rr.Code)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			b := generateHttpTestRequestBody(assert, *test.input)
			req, err := http.NewRequest("POST", "/api/v1/products", b)
			req.Header.Add("authorization", "bearer "+test.headerToken)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// Test_HandleGetProducts
type handleGetProductsTestCase struct {
	name   string
	userId string
	exec   func(*httptest.ResponseRecorder)
}

func Test_HandleGetProducts(t *testing.T) {
	ctx := context.TODO()
	assert, r := handlersTestSetup(ctx, t)

	testCases := []handleGetProductsTestCase{
		{
			name:   "get with valid param",
			userId: uuid.NewString(),
			exec: func(rr *httptest.ResponseRecorder) {
				var rs utils.ResponseMessage[dto.QueryProductsResponseDto]
				err := json.Unmarshal(rr.Body.Bytes(), &rs)
				assert.NotEmpty(rs)
				assert.NoError(err)
				assert.Equal(http.StatusOK, rr.Code)
				assert.Equal(0, rs.Data.Total)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			url := fmt.Sprintf("/api/v1/products/%s", test.userId)
			req, err := http.NewRequest("GET", url, nil)
			assert.NotEmpty(req)
			assert.NoError(err)

			test.exec(executeHttpTestRequest(req, r))
		})
	}
}

// todo:
// User: HandleCheckUserExist
// Product: HandleGetProductById
// Product: HandleUpdateProductById
// Product: HandleDeleteProductById
// Order: HandleCreateOrder
// Order: HandleGetOrders
// Order: HandleGetOrderById
// Order: HandleUpdateOrderById
// Order: HandleDeleteOrderById
// Site: HandleGetSiteUiByUserId
// Site: HandleUpsertSiteUiByUserId
// Album: HandleUploadAlbumImage
// Album: HandleGetAlbumImgs
// Album: HandleUpdateS3ImageDataById
