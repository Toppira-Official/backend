package configs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HttpServer struct {
	envs Environments
}

type HttpServerDeps struct {
	fx.In

	ErrorHandler gin.HandlerFunc `name:"error_handler"`
	Envs         Environments
	Logger       *zap.Logger
}

func NewHttpServer(lc fx.Lifecycle, d HttpServerDeps) *gin.Engine {
	switch d.Envs.MODE.String() {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "develop":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()

	engine.Use(ginzap.Ginzap(d.Logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(d.Logger, true))
	engine.Use(d.ErrorHandler)

	srv := &http.Server{
		Addr:    httpServerPortNumber(d.Envs.PORT.String()),
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				d.Logger.Info("http server started", zap.String("addr", srv.Addr))
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					d.Logger.Fatal("http server crashed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			d.Logger.Info("http server stopping...")
			return srv.Shutdown(ctx)
		},
	})

	return engine
}

func httpServerPortNumber(port string) string {
	return fmt.Sprintf(":%s", port)
}
