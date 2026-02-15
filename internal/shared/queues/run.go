package queues

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RunAsynq(lc fx.Lifecycle, srv *asynq.Server, cl *Client, mux *asynq.ServeMux, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("asynq worker starting...")
			go func() {
				if err := srv.Run(mux); err != nil {
					logger.Error("asynq worker stopped with error", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("asynq worker shutting down...")
			srv.Shutdown()
			return cl.Close()
		},
	})
}
