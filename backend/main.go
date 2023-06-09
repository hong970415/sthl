package main

import (
	"net/http"
	"sthl/api"
	"sthl/config"
	_ "sthl/docs"
	"sthl/logger"
	"sthl/repository"
	"sthl/server"
	"sthl/service"
	"sthl/storage"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server222.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1

func main() {
	fx.New(
		fx.Provide(
			logger.NewDevInfoZapLogger,
			config.NewConfig,

			// db client
			storage.NewPostgresDb,
			storage.NewS3Client,
			// repos
			repository.NewUserRepository,
			repository.NewProductRepository,
			repository.NewOrderRepository,
			repository.NewSiteUiRepository,
			repository.NewImgInfoRepository,

			// services
			service.NewUserService,
			service.NewProductService,
			service.NewOrderService,
			service.NewSiteUiService,
			service.NewAlbumService,

			// http
			api.NewHandler,
			api.NewChiRouter,
			server.NewHttpServer,
		),
		fx.Invoke(
			func(*http.Server) {
			},
		),
	).Run()
}
