package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sthl/config"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHttpServer(lc fx.Lifecycle, lg *zap.Logger, r *chi.Mux, c *config.Config) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.GetPort()),
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			lg.Info("starting HTTP server at", zap.Any("Addr", srv.Addr))
			lg.Info("serving...")

			serve := func() {
				err := srv.Serve(ln)
				if err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						lg.Info(err.Error())
					} else {
						lg.Info("fail to srv.Serve", zap.Error(err))
					}
				}
			}
			go serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
