package configs

import (
	"context"
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

func NewHttpServer(lc fx.Lifecycle, envs Environments, logger *zap.Logger) *gin.Engine {
	engine := gin.New()
	switch envs.MODE.String() {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "develop":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, true))

	srv := &http.Server{
		Addr:    httpServerPortNumber(envs.PORT.String()),
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("http server started", zap.String("addr", srv.Addr))
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("http server crashed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("http server stopping...")
			return srv.Shutdown(ctx)
		},
	})

	return engine
}

func httpServerPortNumber(port string) string {
	return fmt.Sprintf(":%s", port)
}
