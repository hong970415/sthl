package api

import (
	"net/http"
	"os"
	"sthl/constants"
	"sthl/dto"
	"sthl/service"
	"sthl/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Ping godoc
// @Description  ping
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}   utils.ResponseMessage[any]
// @Router       /ping [get]
func HandlePing(w http.ResponseWriter, r *http.Request) {
	// version, exist := os.LookupEnv("VERSION")
	_, exist := os.LookupEnv("VERSION")
	if !exist {
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}
	utils.ResponseSend[any](w, http.StatusOK, "5.2", nil)
	// utils.ResponseSend[any](w, http.StatusOK, version, nil)
}

type IHandler interface {
	// public
	HandleSignup(w http.ResponseWriter, r *http.Request)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCheckUserExist(w http.ResponseWriter, r *http.Request)
	HandleGetProducts(w http.ResponseWriter, r *http.Request)
	HandleGetProductById(w http.ResponseWriter, r *http.Request)
	HandleGetSiteUiByUserId(w http.ResponseWriter, r *http.Request)
	// private
	HandleRefreshAccessToken(w http.ResponseWriter, r *http.Request)
	HandleGetMe(w http.ResponseWriter, r *http.Request)
	HandleUpdateUserPasswordById(w http.ResponseWriter, r *http.Request)
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)
	HandleUpdateProductById(w http.ResponseWriter, r *http.Request)
	HandleDeleteProductById(w http.ResponseWriter, r *http.Request)
	HandleCreateOrder(w http.ResponseWriter, r *http.Request)
	HandleGetOrders(w http.ResponseWriter, r *http.Request)
	HandleGetOrderById(w http.ResponseWriter, r *http.Request)
	HandleUpdateOrderById(w http.ResponseWriter, r *http.Request)
	HandleDeleteOrderById(w http.ResponseWriter, r *http.Request)
	HandleUpsertSiteUiByUserId(w http.ResponseWriter, r *http.Request)
	HandleUploadAlbumImage(w http.ResponseWriter, r *http.Request)
	HandleGetAlbumImgs(w http.ResponseWriter, r *http.Request)
	HandleUpdateS3ImageDataById(w http.ResponseWriter, r *http.Request)

	// for test
	HandleGetUsers(w http.ResponseWriter, r *http.Request)
	HandleGetUserById(w http.ResponseWriter, r *http.Request)
}
type Handler struct {
	logger     *zap.Logger
	userSvc    service.IUserService
	productSvc service.IProductService
	orderSvc   service.IOrderService
	siteuiSvc  service.ISiteUiService
	albumSvc   service.IAlbumService
}

func NewHandler(l *zap.Logger,
	userSvc service.IUserService,
	productSvc service.IProductService,
	orderSvc service.IOrderService,
	siteuiSvc service.ISiteUiService,
	albumSvc service.IAlbumService,
) IHandler {
	return &Handler{
		logger:     l,
		userSvc:    userSvc,
		productSvc: productSvc,
		orderSvc:   orderSvc,
		siteuiSvc:  siteuiSvc,
		albumSvc:   albumSvc,
	}
}

// ****User

// public: HandleSignup
func (h *Handler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract request body
	payload, err := utils.GetRequestBody[dto.CreateUserDto](r.Body)
	if err != nil {
		h.logger.Info("fail to getRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to signup
	result, err := h.userSvc.Signup(ctx, payload)
	if err != nil {
		h.logger.Info("fail to userSvc.Signup", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusCreated, "created", result)
}

// public: HandleLogin
func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract request body
	payload, err := utils.GetRequestBody[dto.LoginDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to login
	result, err := h.userSvc.Login(ctx, payload)
	if err != nil {
		h.logger.Info("fail to userSvc.Login", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}

	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// public: HandleCheckUserExist
func (h *Handler) HandleCheckUserExist(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// get url param
	userIdParam := chi.URLParam(r, "userId")

	// call service to GetUserById
	_, err := h.userSvc.GetUserById(ctx, userIdParam)
	if err != nil {
		h.logger.Info("fail to userSvc.GetUserById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend[any](w, http.StatusOK, "ok", nil)
}

// private: HandleRefreshAccessToken
func (h *Handler) HandleRefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	// extract request body
	payload, err := utils.GetRequestBody[dto.RefreshAccessTokenDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to RefreshAccessToken
	result, err := h.userSvc.RefreshAccessToken(ctx, authenticatedUserInfo, payload)
	if err != nil {
		h.logger.Info("fail to userSvc.RefreshAccessToken", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleGetMe
func (h *Handler) HandleGetMe(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	// call service to GetUserById
	result, err := h.userSvc.GetUserById(ctx, authenticatedUserInfo)
	if err != nil {
		h.logger.Info("fail to userSvc.GetUserById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleUpdateUserPasswordById
func (h *Handler) HandleUpdateUserPasswordById(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	// extract request body
	payload, err := utils.GetRequestBody[dto.UpdateUserPasswordDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to UpdateUserPasswordById
	result, err := h.userSvc.UpdateUserPasswordById(ctx, authenticatedUserInfo, payload)
	if err != nil {
		h.logger.Info("fail to userSvc.UpdateUserPasswordById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleGetUsers
func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract paging
	paging := dto.ExtractPaging(r)
	payload := dto.NewQueryUsersDto(*paging)

	result, err := h.userSvc.GetUsers(ctx, payload)
	if err != nil {
		h.logger.Info("fail to userSvc.GetUsers", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleGetUserById
func (h *Handler) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// get url param
	userIdParam := chi.URLParam(r, "userId")

	// call service to GetUserById
	result, err := h.userSvc.GetUserById(ctx, userIdParam)
	if err != nil {
		h.logger.Info("fail to userSvc.GetUserById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// ****Product

// private: HandleCreateProduct
func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// extract ctx from request
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}
	// extract request body
	payload, err := utils.GetRequestBody[dto.CreateProductDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to signup
	result, err := h.productSvc.CreateProduct(ctx, authenticatedUserInfo, payload)
	if err != nil {
		h.logger.Info("fail to productSvc.CreateProduct", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusCreated, "created", result)
}

// public: HandleGetProducts
func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract paging
	paging := dto.ExtractPaging(r)
	payload := dto.NewQueryProductsDto(*paging)

	// get url param
	userIdParam := chi.URLParam(r, "userId")

	result, err := h.productSvc.GetPrdoucts(ctx, userIdParam, payload)
	if err != nil {
		h.logger.Info("fail to productSvc.GetPrdoucts", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// public: HandleGetProductById
func (h *Handler) HandleGetProductById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// get url param
	// userIdParam := chi.URLParam(r, "userId")
	productIdParam := chi.URLParam(r, "productId")

	result, err := h.productSvc.GetProductById(ctx, productIdParam)
	if err != nil {
		h.logger.Info("fail to productSvc.GetProductById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleUpdateProductById
func (h *Handler) HandleUpdateProductById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// get url param
	userIdParam := chi.URLParam(r, "userId")
	productIdParam := chi.URLParam(r, "productId")

	// extract request body
	payload, err := utils.GetRequestBody[dto.UpdateProductDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	result, err := h.productSvc.UpdateProductById(ctx, userIdParam, productIdParam, payload)
	if err != nil {
		h.logger.Info("fail to productSvc.UpdateProductById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleDeleteProductById
func (h *Handler) HandleDeleteProductById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// get url param
	userIdParam := chi.URLParam(r, "userId")
	productIdParam := chi.URLParam(r, "productId")

	_, err := h.productSvc.SoftDeleteProductById(ctx, userIdParam, productIdParam)
	if err != nil {
		h.logger.Info("fail to productSvc.SoftDeleteProductById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend[any](w, http.StatusOK, "ok", nil)
}

// ****Order

// private: HandleCreateOrder
func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// // extract AccessTokenInfo from ctx
	// authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	// if !ok {
	// 	h.logger.Info("fail to extract userInfo from ctx")
	// 	utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
	// 	return
	// }

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		h.logger.Info("fail to extract userId")
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}

	// extract request body
	payload, err := utils.GetRequestBody[dto.CreateOrderDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	// call service to signup
	result, err := h.orderSvc.CreateOrder(ctx, userId, payload)
	if err != nil {
		h.logger.Info("fail to orderSvc.CreateOrder", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusCreated, "created", result)
}

// private: HandleGetOrders
func (h *Handler) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	// extract paging
	paging := dto.ExtractPaging(r)
	payload := dto.NewQueryOrdersDto(*paging)

	result, err := h.orderSvc.GetOrders(ctx, authenticatedUserInfo, payload)
	if err != nil {
		h.logger.Info("fail to orderSvc.GetOrders", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleGetOrderById
func (h *Handler) HandleGetOrderById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	orderIdParam := chi.URLParam(r, "orderId")
	result, err := h.orderSvc.GetOrderById(ctx, authenticatedUserInfo, orderIdParam)
	if err != nil {
		h.logger.Info("fail to orderSvc.GetOrderById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleUpdateOrderById
func (h *Handler) HandleUpdateOrderById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	orderIdParam := chi.URLParam(r, "orderId")
	// extract request body
	payload, err := utils.GetRequestBody[dto.UpdateOrderDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}
	h.logger.Info("request body", zap.Any("payload", payload))

	result, err := h.orderSvc.UpdateOrderById(ctx, authenticatedUserInfo, orderIdParam, payload)
	if err != nil {
		h.logger.Info("fail to orderSvc.GetOrderById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleDeleteOrderById
func (h *Handler) HandleDeleteOrderById(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}
	orderIdParam := chi.URLParam(r, "orderId")

	result, err := h.orderSvc.SoftDeleteOrderById(ctx, authenticatedUserInfo, orderIdParam)
	if err != nil || !result {
		h.logger.Info("fail to orderSvc.SoftDeleteOrderById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend[any](w, http.StatusOK, "ok", nil)
}

// ****Site

// public: HandleGetSiteUiByUserId
func (h *Handler) HandleGetSiteUiByUserId(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		h.logger.Info("fail to extract userId")
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}

	// call service to GetSiteUiByUserId
	result, err := h.siteuiSvc.GetSiteUiByUserId(ctx, userId)
	if err != nil {
		h.logger.Info("fail to uiSvc.GetSiteUiByUserId", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleUpsertSiteUiByUserId
func (h *Handler) HandleUpsertSiteUiByUserId(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	// extract request body
	payload, err := utils.GetRequestBody[dto.UpsertSiteUiDto](r.Body)
	if err != nil {
		h.logger.Info("fail to GetRequestBody", zap.Error(err))
		utils.ResponseSend[any](w, http.StatusBadRequest, "", nil)
		return
	}

	result, err := h.siteuiSvc.UpsertSiteUiByUserId(ctx, authenticatedUserInfo, payload)
	if err != nil || !result {
		h.logger.Info("fail to siteuiSvc.UpsertSiteUiByUserId", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend[any](w, http.StatusOK, "ok", nil)
}

// ****Album

// private: HandleUploadAlbumImage
func (h *Handler) HandleUploadAlbumImage(w http.ResponseWriter, r *http.Request) {
	result, err := h.albumSvc.UploadFile(r)
	if err != nil {
		h.logger.Info("fail to albumSvc.UploadFiles", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusCreated, "ok", &result)
}

// private: HandleGetAlbumImgs
func (h *Handler) HandleGetAlbumImgs(w http.ResponseWriter, r *http.Request) {
	// get request ctx
	ctx := r.Context()

	// extract paging
	paging := dto.ExtractPaging(r)
	payload := dto.NewQueryImgsInfoDto(*paging)

	// extract AccessTokenInfo from ctx
	authenticatedUserInfo, ok := ctx.Value(constants.AccessTokenInfoKey).(string)
	if !ok {
		h.logger.Info("fail to extract userInfo from ctx")
		utils.ResponseSend[any](w, http.StatusInternalServerError, "", nil)
		return
	}

	result, err := h.albumSvc.GetImgsByUserId(ctx, authenticatedUserInfo, payload)
	if err != nil {
		h.logger.Info("fail to albumSvc.GetImgsByUserId", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", result)
}

// private: HandleUpdateS3ImageDataById
func (h *Handler) HandleUpdateS3ImageDataById(w http.ResponseWriter, r *http.Request) {
	result, err := h.albumSvc.UpdateS3ImageDataById(r)
	if err != nil {
		h.logger.Info("fail to albumSvc.UpdateS3ImageDataById", zap.Error(err))
		utils.HttpErrorResponseSend(w, err)
		return
	}
	utils.ResponseSend(w, http.StatusOK, "ok", &result)
}
