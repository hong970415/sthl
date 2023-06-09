package api

import (
	"fmt"
	"net/http"
	"sthl/authentication"
	"sthl/config"
	"sthl/service"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func NewChiRouter(l *zap.Logger, cfg *config.Config, userSvc service.IUserService, handlers IHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("content-type", "application/json"))

	allowMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowHeaders := []string{
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
		"Origin",
		"Accept",
		"X-Requested-With",
		"X-CSRF-Token",
		"Content-Type",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Credentials",
		"Authorization",
	}
	r.Use(cors.Handler(cors.Options{
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedOrigins:   []string{cfg.GetAllowOrigin()},
		AllowedMethods:   allowMethods,
		AllowedHeaders:   allowHeaders,
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Recoverer)
	RegisterRoute(l, r, userSvc, handlers)
	return r
}

func RegisterRoute(l *zap.Logger, r *chi.Mux, authentor authentication.Authenticator, hdlr IHandler) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4000/swagger/doc.json"), //The url pointing to API definition
	))

	r.Group(func(rt chi.Router) {
		// public
		rt.Get("/api/v1/ping", HandlePing)
		rt.Post("/api/v1/users", hdlr.HandleSignup)
		rt.Post("/api/v1/users/login", hdlr.HandleLogin)
		rt.Get("/api/v1/users/exist/{userId}", hdlr.HandleCheckUserExist)
		rt.Get("/api/v1/products/{userId}", hdlr.HandleGetProducts)
		rt.Get("/api/v1/products/{userId}/{productId}", hdlr.HandleGetProductById)
		rt.Post("/api/v1/orders/{userId}", hdlr.HandleCreateOrder)
		rt.Get("/api/v1/siteui/{userId}", hdlr.HandleGetSiteUiByUserId)
	})
	// private
	r.Group(func(rt chi.Router) {
		rt.Use(authentication.AuthInterceptor(l, authentor))
		rt.Post("/api/v1/users/refreshToken", hdlr.HandleRefreshAccessToken)
		rt.Get("/api/v1/users/me", hdlr.HandleGetMe)
		rt.Put("/api/v1/users/me/pw", hdlr.HandleUpdateUserPasswordById)
		rt.Post("/api/v1/products", hdlr.HandleCreateProduct)
		rt.Put("/api/v1/products/{userId}/{productId}", hdlr.HandleUpdateProductById)
		rt.Delete("/api/v1/products/{userId}/{productId}", hdlr.HandleDeleteProductById)
		rt.Get("/api/v1/orders", hdlr.HandleGetOrders)
		rt.Get("/api/v1/orders/{orderId}", hdlr.HandleGetOrderById)
		rt.Put("/api/v1/orders/{orderId}", hdlr.HandleUpdateOrderById)
		rt.Delete("/api/v1/orders/{orderId}", hdlr.HandleDeleteOrderById)
		rt.Put("/api/v1/siteui", hdlr.HandleUpsertSiteUiByUserId)
		rt.Post("/api/v1/album", hdlr.HandleUploadAlbumImage)
		rt.Get("/api/v1/album", hdlr.HandleGetAlbumImgs)
		rt.Put("/api/v1/album/{imgInfoId}", hdlr.HandleUpdateS3ImageDataById)
		// for test
		rt.Get("/api/v1/users/{userId}", hdlr.HandleGetUserById)
		rt.Get("/api/v1/users", hdlr.HandleGetUsers)
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		value := fmt.Sprintf("method:%s\troute:%s", method, route)
		l.Info(value)
		return nil
	}
	err := chi.Walk(r, walkFunc)
	if err != nil {
		l.Info("fail to walk chi", zap.Error(err))
	}
}
